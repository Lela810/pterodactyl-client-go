package pterodactyl

import "time"

// User -
type User struct {
	ID         int32     `json:"id"`
	ExternalID string    `json:"external_id"`
	UUID       string    `json:"uuid"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Language   string    `json:"language"`
	RootAdmin  bool      `json:"root_admin"`
	Is2FA      bool      `json:"2fa"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (u User) GetEmail() string {
	return u.Email
}
func (u User) GetUsername() string {
	return u.Username
}
func (u User) GetFirstName() string {
	return u.FirstName
}
func (u User) GetLastName() string {
	return u.LastName
}

// Partial User - Only used for creating a new user
type PartialUser struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	// Language  string `json:"language"`
}

func (p PartialUser) GetEmail() string {
	return p.Email
}
func (p PartialUser) GetUsername() string {
	return p.Username
}
func (p PartialUser) GetFirstName() string {
	return p.FirstName
}
func (p PartialUser) GetLastName() string {
	return p.LastName
}

type UsersResponse struct {
	Object string         `json:"object"`
	Data   []UserResponse `json:"data"`
}
type UserResponse struct {
	Object     string `json:"object"`
	Attributes User   `json:"attributes"`
}

// UserInterface - interface for User and PartialUser
type UserInterface interface {
	GetUsername() string
	GetEmail() string
	GetFirstName() string
	GetLastName() string
}

// Node -
type Node struct {
	ID                 int32     `json:"id"`
	UUID               string    `json:"uuid"`
	Public             bool      `json:"public"`
	Name               string    `json:"name"`
	Description        string    `json:"description"`
	LocationID         int32     `json:"location_id"`
	FQDN               string    `json:"fqdn"`
	Scheme             string    `json:"scheme"`
	BehindProxy        bool      `json:"behind_proxy"`
	MaintenanceMode    bool      `json:"maintenance_mode"`
	Memory             int32     `json:"memory"`
	MemoryOverallocate int32     `json:"memory_overallocate"`
	Disk               int32     `json:"disk"`
	DiskOverallocate   int32     `json:"disk_overallocate"`
	UploadSize         int32     `json:"upload_size"`
	DaemonListen       int32     `json:"daemon_listen"`
	DaemonSFTP         int32     `json:"daemon_sftp"`
	DaemonBase         string    `json:"daemon_base"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

func (n Node) GetName() string {
	return n.Name
}
func (n Node) GetDescription() string {
	return n.Description
}
func (n Node) GetPublic() bool {
	return n.Public
}
func (n Node) GetBehindProxy() bool {
	return n.BehindProxy
}
func (n Node) GetMaintenanceMode() bool {
	return n.MaintenanceMode
}
func (n Node) GetLocationID() int32 {
	return n.LocationID
}
func (n Node) GetFQDN() string {
	return n.FQDN
}
func (n Node) GetScheme() string {
	return n.Scheme
}
func (n Node) GetMemory() int32 {
	return n.Memory
}
func (n Node) GetMemoryOverallocate() int32 {
	return n.MemoryOverallocate
}
func (n Node) GetDisk() int32 {
	return n.Disk
}
func (n Node) GetDiskOverallocate() int32 {
	return n.DiskOverallocate
}
func (n Node) GetUploadSize() int32 {
	return n.UploadSize
}
func (n Node) GetDaemonSFTP() int32 {
	return n.DaemonSFTP
}
func (n Node) GetDaemonListen() int32 {
	return n.DaemonListen
}

// Partial Node - Only used for creating a new node
type PartialNode struct {
	Name               string `json:"name"`
	Description        string `json:"description"`
	Public             bool   `json:"public"`
	BehindProxy        bool   `json:"behind_proxy"`
	MaintenanceMode    bool   `json:"maintenance_mode"`
	LocationID         int32  `json:"location_id"`
	FQDN               string `json:"fqdn"`
	Scheme             string `json:"scheme"`
	Memory             int32  `json:"memory"`
	MemoryOverallocate int32  `json:"memory_overallocate"`
	Disk               int32  `json:"disk"`
	DiskOverallocate   int32  `json:"disk_overallocate"`
	UploadSize         int32  `json:"upload_size"`
	DaemonSFTP         int32  `json:"daemon_sftp"`
	DaemonListen       int32  `json:"daemon_listen"`
}

func (n PartialNode) GetName() string {
	return n.Name
}
func (n PartialNode) GetDescription() string {
	return n.Description
}
func (n PartialNode) GetPublic() bool {
	return n.Public
}
func (n PartialNode) GetBehindProxy() bool {
	return n.BehindProxy
}
func (n PartialNode) GetMaintenanceMode() bool {
	return n.MaintenanceMode
}
func (n PartialNode) GetLocationID() int32 {
	return n.LocationID
}
func (n PartialNode) GetFQDN() string {
	return n.FQDN
}
func (n PartialNode) GetScheme() string {
	return n.Scheme
}
func (n PartialNode) GetMemory() int32 {
	return n.Memory
}
func (n PartialNode) GetMemoryOverallocate() int32 {
	return n.MemoryOverallocate
}
func (n PartialNode) GetDisk() int32 {
	return n.Disk
}
func (n PartialNode) GetDiskOverallocate() int32 {
	return n.DiskOverallocate
}
func (n PartialNode) GetUploadSize() int32 {
	return n.UploadSize
}
func (n PartialNode) GetDaemonSFTP() int32 {
	return n.DaemonSFTP
}
func (n PartialNode) GetDaemonListen() int32 {
	return n.DaemonListen
}

type NodesResponse struct {
	Object string         `json:"object"`
	Data   []NodeResponse `json:"data"`
}

type NodeResponse struct {
	Object     string `json:"object"`
	Attributes Node   `json:"attributes"`
}

// Nodes Interface - interface for Node and PartialNode
type NodesInterface interface {
	GetName() string
	GetDescription() string
	GetPublic() bool
	GetBehindProxy() bool
	GetMaintenanceMode() bool
	GetLocationID() int32
	GetFQDN() string
	GetScheme() string
	GetMemory() int32
	GetMemoryOverallocate() int32
	GetDisk() int32
	GetDiskOverallocate() int32
	GetUploadSize() int32
	GetDaemonSFTP() int32
	GetDaemonListen() int32
}

// Node Configuration -
type NodeConfiguration struct {
	Debug   bool   `json:"debug"`
	UUID    string `json:"uuid"`
	TokenID string `json:"token_id"`
	Token   string `json:"token"`
	API     struct {
		Host string `json:"host"`
		Port int32  `json:"port"`
		SSL  struct {
			Enabled bool   `json:"enabled"`
			Cert    string `json:"cert"`
			Key     string `json:"key"`
		} `json:"ssl"`
		UploadLimit int32 `json:"upload_limit"`
	} `json:"api"`
	System struct {
		Data string `json:"data"`
		SFTP struct {
			BindPort int32 `json:"bind_port"`
		} `json:"sftp"`
	} `json:"system"`
	Remote string `json:"remote"`
}

type NodeConfigurationResponse struct {
	Object     string            `json:"object"`
	Attributes NodeConfiguration `json:"attributes"`
}

// Allocation -
type Allocation struct {
	ID       int32  `json:"id"`
	IP       string `json:"ip"`
	Alias    string `json:"alias"`
	Port     int32  `json:"port"`
	Notes    string `json:"notes"`
	Assigned bool   `json:"assigned"`
}

type PartialAllocation struct {
	IP    string `json:"ip"`
	Ports []int  `json:"ports"`
}

type AllocationsResponse struct {
	Object string               `json:"object"`
	Data   []AllocationResponse `json:"data"`
}

type AllocationResponse struct {
	Object     string     `json:"object"`
	Attributes Allocation `json:"attributes"`
}

// Location -
type Location struct {
	ID        int32     `json:"id"`
	Short     string    `json:"short"`
	Long      string    `json:"long"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

// ServerDatabase -
type ServerDatabase struct {
	ID             int32     `json:"id"`
	Server         int32     `json:"server"`
	Host           int32     `json:"host"`
	Database       string    `json:"database"`
	Username       string    `json:"username"`
	Remote         string    `json:"remote"`
	MaxConnections int32     `json:"max_connections"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Relationships  struct {
		Password struct {
			Object     string `json:"object"`
			Attributes struct {
				Password string `json:"password"`
			} `json:"attributes"`
		} `json:"password"`
		Host struct {
			Object     string `json:"object"`
			Attributes []Host `json:"attributes"`
		} `json:"host"`
	} `json:"relationships"`
}

// Host -
type Host struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	Host      string    `json:"host"`
	Port      int32     `json:"port"`
	Username  string    `json:"username"`
	Node      int32     `json:"node"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Server -
type Server struct {
	ID            int32     `json:"id"`
	ExternalID    string    `json:"external_id"`
	UUID          string    `json:"uuid"`
	Identifier    string    `json:"identifier"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Suspended     bool      `json:"suspended"`
	Limits        Limits    `json:"limits"`
	FeatureLimit  Feature   `json:"feature_limits"`
	User          int32     `json:"user"`
	Node          int32     `json:"node"`
	Allocation    int32     `json:"allocation"`
	Nest          int32     `json:"nest"`
	Egg           int32     `json:"egg"`
	Pack          int32     `json:"pack"`
	Container     Container `json:"container"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Relationships struct {
		Databases struct {
			Object string           `json:"object"`
			Data   []ServerDatabase `json:"data"`
		} `json:"databases"`
	} `json:"relationships"`
}

// Limits -
type Limits struct {
	Memory  int32 `json:"memory"`
	Swap    int32 `json:"swap"`
	Disk    int32 `json:"disk"`
	IO      int32 `json:"io"`
	CPU     int32 `json:"cpu"`
	Threads int32 `json:"threads"`
}

// Feature -
type Feature struct {
	Databases   int32 `json:"databases"`
	Allocations int32 `json:"allocations"`
	Backups     int32 `json:"backups"`
}

// Container -
type Container struct {
	StartupCommand string `json:"startup_command"`
	Image          string `json:"image"`
	Installed      bool   `json:"installed"`
	Environment    struct {
		ServerJarfile string `json:"SERVER_JARFILE"`
		VanillaVer    string `json:"VANILLA_VERSION"`
		Startup       string `json:"STARTUP"`
		PServerLoc    string `json:"P_SERVER_LOCATION"`
		PServerUUID   string `json:"P_SERVER_UUID"`
	} `json:"environment"`
}

// Egg -
type Egg struct {
	Object     string `json:"object"`
	Attributes struct {
		ID            int32     `json:"id"`
		UUID          string    `json:"uuid"`
		Name          string    `json:"name"`
		Nest          int32     `json:"nest"`
		Author        string    `json:"author"`
		Description   string    `json:"description"`
		DockerImage   string    `json:"docker_image"`
		Config        EggConfig `json:"config"`
		Startup       string    `json:"startup"`
		Script        EggScript `json:"script"`
		CreatedAt     time.Time `json:"created_at"`
		UpdatedAt     time.Time `json:"updated_at"`
		Relationships struct {
			Nest struct {
				Object     string `json:"object"`
				Attributes struct {
					ID          int32     `json:"id"`
					UUID        string    `json:"uuid"`
					Author      string    `json:"author"`
					Name        string    `json:"name"`
					Description string    `json:"description"`
					CreatedAt   time.Time `json:"created_at"`
					UpdatedAt   time.Time `json:"updated_at"`
				} `json:"attributes"`
			} `json:"nest"`
			Servers struct {
				Object string        `json:"object"`
				Data   []interface{} `json:"data"`
			} `json:"servers"`
		} `json:"relationships"`
	} `json:"attributes"`
}

// EggConfig -
type EggConfig struct {
	Files   map[string]EggConfigFile `json:"files"`
	Startup struct {
		Done            string   `json:"done"`
		UserInteraction []string `json:"userInteraction"`
	} `json:"startup"`
	Stop string `json:"stop"`
	Logs struct {
		Custom   bool   `json:"custom"`
		Location string `json:"location"`
	} `json:"logs"`
	Extends interface{} `json:"extends"`
}

// EggConfigFile -
type EggConfigFile struct {
	Parser string                 `json:"parser"`
	Find   map[string]interface{} `json:"find"`
}

// EggScript -
type EggScript struct {
	Privileged bool        `json:"privileged"`
	Install    string      `json:"install"`
	Entry      string      `json:"entry"`
	Container  string      `json:"container"`
	Extends    interface{} `json:"extends"`
}

// Nest -
type Nest struct {
	ID          int32     `json:"id"`
	UUID        string    `json:"uuid"`
	Author      string    `json:"author"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
