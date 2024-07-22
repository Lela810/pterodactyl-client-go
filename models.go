package pterodactyl

import "time"

// User -
type User struct {
	ID         int       `json:"id"`
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

// Node -
type Node struct {
	ID                 int       `json:"id"`
	UUID               string    `json:"uuid"`
	Public             bool      `json:"public"`
	Name               string    `json:"name"`
	Description        string    `json:"description"`
	LocationID         int       `json:"location_id"`
	FQDN               string    `json:"fqdn"`
	Scheme             string    `json:"scheme"`
	BehindProxy        bool      `json:"behind_proxy"`
	MaintenanceMode    bool      `json:"maintenance_mode"`
	Memory             int       `json:"memory"`
	MemoryOverallocate int       `json:"memory_overallocate"`
	Disk               int       `json:"disk"`
	DiskOverallocate   int       `json:"disk_overallocate"`
	UploadSize         int       `json:"upload_size"`
	DaemonListen       int       `json:"daemon_listen"`
	DaemonSFTP         int       `json:"daemon_sftp"`
	DaemonBase         string    `json:"daemon_base"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

// Location -
type Location struct {
	ID        int       `json:"id"`
	Short     string    `json:"short"`
	Long      string    `json:"long"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

// ServerDatabase -
type ServerDatabase struct {
	ID             int       `json:"id"`
	Server         int       `json:"server"`
	Host           int       `json:"host"`
	Database       string    `json:"database"`
	Username       string    `json:"username"`
	Remote         string    `json:"remote"`
	MaxConnections int       `json:"max_connections"`
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
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Host      string    `json:"host"`
	Port      int       `json:"port"`
	Username  string    `json:"username"`
	Node      int       `json:"node"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Server -
type Server struct {
	ID            int       `json:"id"`
	ExternalID    string    `json:"external_id"`
	UUID          string    `json:"uuid"`
	Identifier    string    `json:"identifier"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Suspended     bool      `json:"suspended"`
	Limits        Limits    `json:"limits"`
	FeatureLimit  Feature   `json:"feature_limits"`
	User          int       `json:"user"`
	Node          int       `json:"node"`
	Allocation    int       `json:"allocation"`
	Nest          int       `json:"nest"`
	Egg           int       `json:"egg"`
	Pack          int       `json:"pack"`
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
	Memory  int `json:"memory"`
	Swap    int `json:"swap"`
	Disk    int `json:"disk"`
	IO      int `json:"io"`
	CPU     int `json:"cpu"`
	Threads int `json:"threads"`
}

// Feature -
type Feature struct {
	Databases   int `json:"databases"`
	Allocations int `json:"allocations"`
	Backups     int `json:"backups"`
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
		ID            int       `json:"id"`
		UUID          string    `json:"uuid"`
		Name          string    `json:"name"`
		Nest          int       `json:"nest"`
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
					ID          int       `json:"id"`
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
	ID          int       `json:"id"`
	UUID        string    `json:"uuid"`
	Author      string    `json:"author"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Response -
type Response struct {
	Object     string `json:"object"`
	Attributes User   `json:"attributes"`
}
