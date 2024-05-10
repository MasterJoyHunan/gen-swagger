package generator

import "github.com/zeromicro/go-zero/tools/goctl/api/spec"

// 将嵌套结构体的一维化
func deconstructionMember(d spec.DefineStruct) []spec.Member {
	var members []spec.Member
	for _, member := range d.Members {
		if member.IsInline { // TODO 只支持内嵌一层，如果嵌套多层将无法解析
			members = append(members, deconstructionMember(member.Type.(spec.DefineStruct))...)
		} else {
			members = append(members, member)
		}
	}
	return members
}
