#!/usr/bin/env python3
"""
自动移除点导入并更新类型引用的脚本
用于优化Go项目的编译内存消耗
"""

import os
import re
import sys
from pathlib import Path
from typing import List, Tuple, Set

# 禁用输出缓冲，确保print立即显示
sys.stdout.reconfigure(line_buffering=True)

class DotImportRemover:
    def __init__(self, project_root: str):
        self.project_root = Path(project_root)
        self.api_dir = self.project_root / "api"
        self.models_dir = self.project_root / "models"
        self.modified_files = []
        self.type_cache = {}
        
        # 预编译正则表达式
        self.dot_import_pattern = re.compile(r'\.\s*"github\.com/oceanengine/ad_open_sdk_go/models"')
        self.normal_import_pattern = re.compile(r'"github\.com/oceanengine/ad_open_sdk_go/models"')
        self.import_block_pattern = re.compile(r'(import\s*\([^)]*\))', re.DOTALL)
        self.package_pattern = re.compile(r'^(package\s+\w+)', re.MULTILINE)
        self.models_prefix_pattern = re.compile(r'models\.models\.')
        
    def collect_model_types(self) -> Set[str]:
        """收集所有模型类型名称"""
        types = set()
        
        if not self.models_dir.exists():
            print(f"警告: models目录不存在: {self.models_dir}", flush=True)
            return types
            
        for model_file in self.models_dir.glob("*.go"):
            try:
                with open(model_file, 'r', encoding='utf-8') as f:
                    content = f.read()
                    
                # 匹配类型定义: type TypeName struct/interface/string/基本类型
                type_pattern = r'^type\s+(\w+)\s+(?:struct|interface|string|int\d*|float\d*|bool|byte|rune|uint\d*|complex\d*|error)\b'
                matches = re.findall(type_pattern, content, re.MULTILINE)
                types.update(matches)
                
                # 匹配常量定义: const ( ... )
                const_pattern = r'^const\s+\(\s*\n((?:\s+\w+\s+\w+\s*=\s*"[^"]*"\s*\n)+)\s*\)'
                for match in re.finditer(const_pattern, content, re.MULTILINE):
                    const_block = match.group(1)
                    const_names = re.findall(r'\s+(\w+)\s+\w+\s*=', const_block)
                    types.update(const_names)
                    
            except Exception as e:
                print(f"警告: 读取文件失败 {model_file}: {e}", flush=True)
                
        print(f"收集到 {len(types)} 个模型类型", flush=True)
        return types
    
    def remove_dot_import(self, file_path: Path, model_types: Set[str]) -> bool:
        """移除点导入并更新类型引用"""
        try:
            with open(file_path, 'r', encoding='utf-8') as f:
                content = f.read()
            
            original_content = content
            has_dot_import = False
            
            # 收集当前文件中定义的类型
            local_types = set()
            # 匹配类型定义：type TypeName Alias 或 type TypeName struct/interface/string/基本类型
            # 使用 findall 匹配所有组，收集所有类型名
            type_pattern = r'^type\s+(\w+)\s+(\w+)'
            for match in re.finditer(type_pattern, content, re.MULTILINE):
                # 添加类型名（第一个词）
                local_types.add(match.group(1))
            
            # 检查并移除点导入
            if self.dot_import_pattern.search(content):
                has_dot_import = True
                
                # 移除点导入行
                lines = content.split('\n')
                new_lines = []
                for line in lines:
                    if not self.dot_import_pattern.search(line):
                        new_lines.append(line)
                
                content = '\n'.join(new_lines)
                
                # 添加正常导入（如果还没有）
                if not self.normal_import_pattern.search(content):
                    # 找到import块并添加
                    match = self.import_block_pattern.search(content)
                    if match:
                        import_block = match.group(1)
                        # 在import块末尾添加（在 ) 之前）
                        new_import = import_block.rstrip(')') + '\n\t"github.com/oceanengine/ad_open_sdk_go/models"\n)'
                        content = content.replace(import_block, new_import)
                    else:
                        # 创建新的import块
                        content = self.package_pattern.sub(
                            r'\1\n\nimport (\n\t"github.com/oceanengine/ad_open_sdk_go/models"\n)',
                            content,
                            count=1
                        )
            
            # 更新类型引用（添加models前缀）
            # 只有当文件中有点导入时才进行替换
            if has_dot_import:
                # 只替换从 models 包导入的类型，不替换当前文件中定义的类型
                types_to_replace = model_types - local_types
                
                if types_to_replace:
                    # 先找出文件中实际使用的类型
                    used_types = set()
                    for type_name in types_to_replace:
                        if type_name in content:
                            used_types.add(type_name)
                    
                    # 只对实际使用的类型进行替换
                    if used_types:
                        # 按长度降序排序，避免部分匹配
                        sorted_types = sorted(used_types, key=len, reverse=True)
                        
                        # 构建一个大的正则表达式，一次性匹配所有类型
                        type_pattern = r'\b(' + '|'.join(re.escape(t) for t in sorted_types) + r')\b'
                        type_regex = re.compile(type_pattern)
                        
                        # 使用替换函数，根据上下文决定是否添加 models. 前缀
                        original_content_for_check = content
                        
                        def add_models_prefix(match):
                            type_name = match.group(1)
                            start_pos = match.start()
                            end_pos = match.end()
                            
                            # 检查前面是否已经有 models. 前缀
                            if start_pos >= 7 and original_content_for_check[start_pos-7:start_pos] == 'models.':
                                return type_name
                            
                            # 获取当前行的内容（使用原始内容）
                            line_start = original_content_for_check.rfind('\n', 0, start_pos) + 1
                            line_end = original_content_for_check.find('\n', end_pos)
                            if line_end == -1:
                                line_end = len(original_content_for_check)
                            line_content = original_content_for_check[line_start:line_end]
                            before_match = line_content[:start_pos-line_start]
                            
                            # 检查是否在字符串中
                            if before_match.count('"') % 2 == 1:
                                return type_name
                            
                            # 检查是否在注释中
                            if '//' in before_match:
                                return type_name
                            
                            # 检查是否是函数/方法名
                            full_before = before_match + type_name
                            after_char = original_content_for_check[end_pos:end_pos+1] if end_pos < len(original_content_for_check) else ''
                            
                            # 模式1: func (receiver) MethodName( - 方法名
                            if re.search(r'func\s*\([^)]*\)\s+\w+\s*$', full_before) and after_char == '(':
                                return type_name
                            
                            # 模式2: func FunctionName( - 函数名
                            if re.search(r'func\s+\w+\s*$', full_before) and after_char == '(':
                                return type_name
                            
                            # 检查是否是方法调用：object.Method(
                            if start_pos > 0 and original_content_for_check[start_pos-1] == '.' and after_char == '(':
                                return type_name
                            
                            # 检查是否是类型转换：Type(value)
                            if start_pos > 0 and original_content_for_check[start_pos-1] == '(' and after_char == ')':
                                return f'models.{type_name}'
                            
                            # 检查是否是变量声明：var name Type 或 name Type 或 name := Type
                            if re.search(r'\bvar\b', before_match):
                                return f'models.{type_name}'
                            
                            # 检查是否是结构体字段：FieldName Type
                            if re.search(r'\w+\s+$', before_match):
                                return f'models.{type_name}'
                            
                            # 检查是否是函数参数或返回值：(param Type) 或 () Type
                            if start_pos > 0 and original_content_for_check[start_pos-1] in '(,*':
                                return f'models.{type_name}'
                            
                            # 检查是否是数组/切片/指针类型：[]Type, *Type, map[string]Type
                            if start_pos > 0 and original_content_for_check[start_pos-1] in '[]*':
                                return f'models.{type_name}'
                            
                            # 默认情况下添加 models. 前缀
                            return f'models.{type_name}'
                        
                        content = type_regex.sub(add_models_prefix, content)
                
                # 移除重复的models.前缀
                content = self.models_prefix_pattern.sub('models.', content)
            
            # 只在有修改时写入文件
            if content != original_content:
                with open(file_path, 'w', encoding='utf-8') as f:
                    f.write(content)
                self.modified_files.append(str(file_path.relative_to(self.project_root)))
                return True
                
            return False
            
        except Exception as e:
            print(f"错误: 处理文件失败 {file_path}: {e}", flush=True)
            return False
    
    def process_api_files(self) -> int:
        """处理所有API文件"""
        if not self.api_dir.exists():
            print(f"错误: api目录不存在: {self.api_dir}", flush=True)
            return 0
            
        print("正在收集模型类型...", flush=True)
        model_types = self.collect_model_types()
        
        if not model_types:
            print("警告: 未找到任何模型类型", flush=True)
            return 0
            
        print(f"\n开始处理API文件...", flush=True)
        
        api_files = list(self.api_dir.glob("*.go"))
        processed_count = 0
        
        for api_file in api_files:
            if api_file.name == "client.go":
                continue  # 跳过client.go，单独处理
                
            if self.remove_dot_import(api_file, model_types):
                processed_count += 1
                print(f"✓ {api_file.name}", flush=True)
        
        return processed_count
    
    def process_client_file(self) -> bool:
        """处理client.go文件"""
        client_file = self.api_dir / "client.go"
        if not client_file.exists():
            return False
            
        try:
            with open(client_file, 'r', encoding='utf-8') as f:
                content = f.read()
            
            original_content = content
            
            # 检查是否有点导入
            has_dot_import = bool(self.dot_import_pattern.search(content))
            
            if not has_dot_import:
                return False
            
            # 收集当前文件中定义的类型
            local_types = set()
            type_pattern = r'^type\s+(\w+)\s+(\w+)'
            for match in re.finditer(type_pattern, content, re.MULTILINE):
                local_types.add(match.group(1))
            
            # 收集需要替换的类型
            types_to_replace = set()
            for type_name in self.type_cache:
                if type_name in content and type_name not in local_types:
                    types_to_replace.add(type_name)
            
            # 移除点导入行
            lines = content.split('\n')
            new_lines = []
            for line in lines:
                if not self.dot_import_pattern.search(line):
                    new_lines.append(line)
            
            content = '\n'.join(new_lines)
            
            # 添加正常导入（如果还没有）
            if not self.normal_import_pattern.search(content):
                match = self.import_block_pattern.search(content)
                if match:
                    import_block = match.group(1)
                    new_import = import_block.rstrip(')') + '\n\t"github.com/oceanengine/ad_open_sdk_go/models"\n)'
                    content = content.replace(import_block, new_import)
            
            # 替换类型引用
            if types_to_replace:
                sorted_types = sorted(types_to_replace, key=len, reverse=True)
                type_pattern = r'\b(' + '|'.join(re.escape(t) for t in sorted_types) + r')\b'
                type_regex = re.compile(type_pattern)
                
                original_for_check = content
                
                def add_prefix(match):
                    type_name = match.group(1)
                    start_pos = match.start()
                    
                    if start_pos >= 7 and original_for_check[start_pos-7:start_pos] == 'models.':
                        return type_name
                    
                    line_start = original_for_check.rfind('\n', 0, start_pos) + 1
                    line_end = original_for_check.find('\n', start_pos)
                    if line_end == -1:
                        line_end = len(original_for_check)
                    line_content = original_for_check[line_start:line_end]
                    before = line_content[:start_pos-line_start]
                    
                    if before.count('"') % 2 == 1:
                        return type_name
                    if '//' in before:
                        return type_name
                    
                    full_before = before + type_name
                    after = original_for_check[match.end():match.end()+1] if match.end() < len(original_for_check) else ''
                    
                    if re.search(r'func\s*\([^)]*\)\s+\w+\s*$', full_before) and after == '(':
                        return type_name
                    if re.search(r'func\s+\w+\s*$', full_before) and after == '(':
                        return type_name
                    if start_pos > 0 and original_for_check[start_pos-1] == '.' and after == '(':
                        return type_name
                    if start_pos > 0 and original_for_check[start_pos-1] == '(' and after == ')':
                        return f'models.{type_name}'
                    if re.search(r'\bvar\b', before):
                        return f'models.{type_name}'
                    if re.search(r'\w+\s+$', before):
                        return f'models.{type_name}'
                    if start_pos > 0 and original_for_check[start_pos-1] in '(,*':
                        return f'models.{type_name}'
                    if start_pos > 0 and original_for_check[start_pos-1] in '[]*':
                        return f'models.{type_name}'
                    
                    return f'models.{type_name}'
                
                content = type_regex.sub(add_prefix, content)
                content = self.models_prefix_pattern.sub('models.', content)
            
            if content != original_content:
                with open(client_file, 'w', encoding='utf-8') as f:
                    f.write(content)
                self.modified_files.append("api/client.go")
                return True
                
            return False
            
        except Exception as e:
            print(f"错误: 处理client.go失败: {e}", flush=True)
            return False
    
    def verify_compilation(self) -> bool:
        """验证代码是否可以编译"""
        print("\n正在验证编译...", flush=True)
        import subprocess
        
        try:
            result = subprocess.run(
                ['go', 'build', './...'],
                cwd=self.project_root,
                capture_output=True,
                text=True,
                timeout=300
            )
            
            if result.returncode == 0:
                print("✓ 编译成功", flush=True)
                return True
            else:
                print("✗ 编译失败:", flush=True)
                print(result.stderr, flush=True)
                return False
                
        except subprocess.TimeoutExpired:
            print("✗ 编译超时", flush=True)
            return False
        except Exception as e:
            print(f"✗ 编译验证失败: {e}", flush=True)
            return False
    
    def run(self, verify: bool = True) -> bool:
        """执行完整的优化流程"""
        print("=" * 60, flush=True)
        print("开始移除点导入优化", flush=True)
        print("=" * 60, flush=True)
        
        # 收集模型类型并缓存
        print("正在收集模型类型...", flush=True)
        self.type_cache = self.collect_model_types()
        
        # 处理API文件
        api_count = self.process_api_files()
        
        # 处理client.go
        client_modified = self.process_client_file()
        
        print(f"\n处理完成:", flush=True)
        print(f"  - 修改的API文件: {api_count}", flush=True)
        print(f"  - 修改client.go: {'是' if client_modified else '否'}", flush=True)
        print(f"  - 总计修改文件: {len(self.modified_files)}", flush=True)
        
        if self.modified_files:
            print(f"\n修改的文件列表:", flush=True)
            for f in self.modified_files:
                print(f"  - {f}", flush=True)
        
        # 验证编译
        if verify and self.modified_files:
            if not self.verify_compilation():
                print("\n⚠️  警告: 编译验证失败，请检查修改", flush=True)
                return False
        
        print("\n" + "=" * 60, flush=True)
        print("优化完成!", flush=True)
        print("=" * 60, flush=True)
        
        return len(self.modified_files) > 0


def main():
    import argparse
    
    print("🚀 优化脚本启动...", flush=True)
    
    parser = argparse.ArgumentParser(description='移除Go项目中的点导入以优化编译内存')
    parser.add_argument('--project-root', '-p', 
                       default='.',
                       help='项目根目录路径（默认: 当前目录）')
    parser.add_argument('--no-verify', '-n',
                       action='store_true',
                       help='跳过编译验证')
    
    args = parser.parse_args()
    
    print(f"📁 项目根目录: {args.project_root}", flush=True)
    print(f"🔍 跳过验证: {args.no_verify}", flush=True)
    
    project_root = Path(args.project_root).resolve()
    
    if not (project_root / "go.mod").exists():
        print(f"❌ 错误: {project_root} 不是有效的Go项目（未找到go.mod）", flush=True)
        sys.exit(1)
    
    print(f"✅ 找到 go.mod 文件", flush=True)
    
    remover = DotImportRemover(str(project_root))
    success = remover.run(verify=not args.no_verify)
    
    print(f"🏁 脚本执行完成，结果: {'成功' if success else '失败'}", flush=True)
    
    sys.exit(0 if success else 1)


if __name__ == '__main__':
    main()
