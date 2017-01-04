package maestropanel

type DomainExecutionResult struct {
    Result
    Details DomainOperationResult
}

type GetDomainAliasResult struct {
    Result
    Details []string
}

type DomainOperationResult struct {
    Code int32
    Message string
    Id int32
    Name string
    Username string
    Password string
    DomainUser bool
    IpString string
    ModuleResults []DomainOperationModuleResult
}

type DomainOperationModuleResult struct {
    Status bool
    Msg string
    Name string
    PType string
}

type DomainListItemResult struct {
    Result
    Details DomainListItem
}

type LimitResult struct {
    Result
    Details []Limit
}

type Limit struct {
    Name string
    ModuleName string
    FriendlyName string
    Limit int32
    Usage int32
    Group string
}

type Domain struct {
    Name string `json:"name"`
    PlanAlias string `json:"planAlias"`
    UserName string `json:"username"`
    Password string `json:"password"`
    ActiveDomainUser bool `json:"activedomainuser"`
    FirstName string `json:"firstname"`
    LastName string `json:"lastname"`
    Email string `json:"email"`
    Expiration string `json:"expiration"`
    IPAddr string `json:"ipaddr"`
}

type SubDomain struct {
    DomainName string `json:"name"`
    SubDomain string `json:"subdomain"`
    FTPUser string `json:"ftpuser"`
}

type DomainPassword struct {
    Name string `json:"name"`
    NewPassword string `json:"newpassword"`
}

type DomainAlias struct {
    Name string `json:"name"`
    Alias string `json:"alias"`
}

type SetSubDomainFTPAccount struct {
    Name string `json:"name"`
    SubDomain string `json:"subdomain"`
    NewFTPUser string `json:"newftpuser"`
}

type GetListResult struct {
	Result
	Details []DomainListItem 
}

type GetSubDomainsResult struct {
    Result
    Details []SubDomainConfig
}

type SubDomainConfig struct {
    Name string
    FtpUser string
}

type DomainListItem struct {
    Id int32
    Name string
    ExpirationDate string
    Status int32
    OwnerName string
    Email int32
    Disk int32
}

type ChangeIpAddres struct {
    DomainName string `json:"name"`
    NewIPAddres string `json:"newip"`
}

type Forward struct {
    DomainName string `json:"name"`
    Enabled bool `json:"enabled"`
    Destination string `json:"destination"`
    ExacDestination bool `json:"exacDestination"`
    ChildOnly bool `json:"childOnly"`
    StatusCode int `json:"statusCode"`
}

type Reseller struct {
    DomainName string `json:"name"`
    ResellerName string `json:"resellerName"`
}

type DomainPlan struct {
    DomainName string `json:"name"`
    PlanAlias string `json:"planAlias"`
    Action string `json:"action"`
}

type NETRuntimeVersion struct {
    DomainName string `json:"name"`
    RuntimeVersion string `json:"runtimeVersion"`
}

type NETRuntimeInfo struct {
    Name string
    Runtime string
    Mode string
}

type NETRuntimeResult struct {
    Result
    Details NETRuntimeInfo
}

type Path struct {
    DomainName string `json:"name"`
    Path string `json:"path"`
}

type DiskItem struct {
    Name string
    Tpy string
    Read bool
    Write bool
    IsLock bool
    CreateDate string
    ModifyDate string
    Size int64
}

type GetItemsResult struct {
    Result
    Details []DiskItem
}

type Zip struct {
    DomainName string `json:"name"`
    ZipFilePath string `json:"zipFilePath"`
    Items []string
}

type ProtectStatsArea struct {
    DomainName string `json:"name"`
    Username string `json:"username"`
    Password string `json:"password"`
}

type MailBox struct {
    Account string `json:"account"`
    Password string `json:"password"`
    Quota int32 `json:"quota"`
    Redirect bool `json:"redirect"`
    Remail string `json:"remail"`
    DisplayName string `json:"displayname"`
}

type GetMailListResult struct {
    Result
    Details ExportPostOffice
}

type ExportPostOffice struct {
    Name string
    Quota int32
    Accounts []Account
}

type Account struct {
    Name string
    Status bool
    Quota int32
    Usage int32
}

type DatabaseInfo struct {
    DBType string `json:"dbtype"`
    Database string `json:"database"`
    Username string `json:"username"`
    Password string `json:"password"`
    Quota int32 `json:"quota"`
    Host string `json:"host"`
}

type DatabaseUserPasswordChangeModel struct {
    DomainName string `json:"name"`
    DBType string `json:"dbtype"`
    Database string `json:"database"`
    Username string `json:"username"`
    Password string `json:"newpassword"`
}

type GetDatabaseListResult struct {
    Result
    Details []SqlConfig
}

type SqlConfig struct {
    Name string
    DiskQuota int32
    DiskUsage int32
    Collation string
    DbType string
    Users []DbUser
}

type DbUser struct {
    Username string
    Password string
    Host string
    Rights string
}

type DatabaseUserPermission struct {
    DomainName string `json:"name"`
    DBType string `json:"dbtype"`
    Database string `json:"database"`
    Username string `json:"username"`
    Permissions string `json:"permissions"`
}