package lib

const (
	DRIVER_NAME string = "mysql"
)

type ROLE int32

const (
	ROLE_PLAYER ROLE = 0
	ROLE_ADMIN  ROLE = 1
)

// Enum value maps for ROLE.
// var (
// 	ROLE_name = map[int32]string{
// 		0: "PLAYER",
// 		1: "ADMIN",
// 	}
// 	ROLE_value = map[string]int32{
// 		"PLAYER": 0,
// 		"ADMIN":  1,
// 	}
// )

// func (x ROLE) Enum() *ROLE {
// 	p := new(ROLE)
// 	*p = x
// 	return p
// }

// func (x ROLE) String() string {
// 	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
// }

// func (ROLE) Descriptor() protoreflect.EnumDescriptor {
// 	return file_const_proto_enumTypes[2].Descriptor()
// }

// func (ROLE) Type() protoreflect.EnumType {
// 	return &file_const_proto_enumTypes[2]
// }

// func (x ROLE) Number() protoreflect.EnumNumber {
// 	return protoreflect.EnumNumber(x)
// }

// // Deprecated: Use ROLE.Descriptor instead.
// func (ROLE) EnumDescriptor() ([]byte, []int) {
// 	return file_const_proto_rawDescGZIP(), []int{2}
// }
