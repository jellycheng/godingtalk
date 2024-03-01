package godingtalk

type DepartmentReqDto struct {
	Language string `json:"language,omitempty"`
	DeptID   int    `json:"dept_id,omitempty"`
}

type DepartmentSimpleInfo struct {
	AutoAddUser     bool   `json:"auto_add_user"`     // 新人是否自动加入该群，true是，false否
	CreateDeptGroup bool   `json:"create_dept_group"` //是否同步创建一个关联此部门的企业群，true是，false否
	DeptID          int    `json:"dept_id"`           //部门ID
	Name            string `json:"name"`              //部门名称
	ParentID        int    `json:"parent_id"`         //上级部门
}

type DepartmentListRespDto struct {
	Errcode   int                    `json:"errcode"`
	Errmsg    string                 `json:"errmsg"`
	RequestID string                 `json:"request_id"`
	Result    []DepartmentSimpleInfo `json:"result"`
}

// 部门树
type DepartmentSimpleTree struct {
	AutoAddUser     bool                   `json:"auto_add_user"`     // 新人是否自动加入该群，true是，false否
	CreateDeptGroup bool                   `json:"create_dept_group"` //是否同步创建一个关联此部门的企业群，true是，false否
	DeptID          int                    `json:"dept_id"`           //部门ID
	Name            string                 `json:"name"`              //部门名称
	ParentID        int                    `json:"parent_id"`         //上级部门
	SubDeptList     []DepartmentSimpleTree `json:"sub_dept_list"`     // 子部门
}

type DeptUseridDto struct {
	UseridList []string `json:"userid_list"`
}

type DeptUseridListRespDto struct {
	Errcode   int           `json:"errcode"`
	Errmsg    string        `json:"errmsg"`
	RequestID string        `json:"request_id"`
	Result    DeptUseridDto `json:"result"`
}

type UserProfileDto struct {
	Active        bool   `json:"active"`
	Admin         bool   `json:"admin"`
	Avatar        string `json:"avatar"`
	Boss          bool   `json:"boss"`
	DeptIDList    []int  `json:"dept_id_list"`
	DeptOrderList []struct {
		DeptID int   `json:"dept_id"`
		Order  int64 `json:"order"`
	} `json:"dept_order_list"`
	ExclusiveAccount bool `json:"exclusive_account"`
	HideMobile       bool `json:"hide_mobile"`
	LeaderInDept     []struct {
		DeptID int  `json:"dept_id"`
		Leader bool `json:"leader"`
	} `json:"leader_in_dept"`
	Mobile     string `json:"mobile"`
	Name       string `json:"name"`
	RealAuthed bool   `json:"real_authed"`
	RoleList   []struct {
		GroupName string `json:"group_name"`
		ID        int    `json:"id"`
		Name      string `json:"name"`
	} `json:"role_list"`
	Senior    bool   `json:"senior"`
	StateCode string `json:"state_code"`
	Unionid   string `json:"unionid"`
	Userid    string `json:"userid"`
}

type UserProfileRespDto struct {
	Errcode   int            `json:"errcode"`
	Errmsg    string         `json:"errmsg"`
	RequestID string         `json:"request_id"`
	Result    UserProfileDto `json:"result"`
}
