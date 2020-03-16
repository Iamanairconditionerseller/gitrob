package matching

import "regexp"

type FileSignatureType struct {
	Extension string
	Filename  string
	Path      string
}

var fileSignatureTypes = FileSignatureType{
	Extension: "extension",
	Filename:  "filename",
	Path:      "path",
}

var FileSignatures = []FileSignature{
	SimpleFileSignature{
		Part:        fileSignatureTypes.Extension,
		MatchOn:     ".pem",
		Description: "Potential cryptographic private key",
		Comment:     "",
	},
	SimpleFileSignature{
		Part:        fileSignatureTypes.Extension,
		MatchOn:     ".log",
		Description: "Log file",
		Comment:     "Log files can contain secret HTTP endpoints, session IDs, API keys and other goodies",
	},
	SimpleFileSignature{
		Part:        fileSignatureTypes.Extension,
		MatchOn:     ".pkcs12",
		Description: "Potential cryptographic key bundle",
		Comment:     "",
	},
	SimpleFileSignature{
		Part:        fileSignatureTypes.Extension,
		MatchOn:     ".p12",
		Description: "Potential cryptographic key bundle",
		Comment:     "",
	},
	SimpleFileSignature{
		Part:        fileSignatureTypes.Extension,
		MatchOn:     ".pfx",
		Description: "Potential cryptographic key bundle",
		Comment:     "",
	},
	SimpleFileSignature{
		Part:        fileSignatureTypes.Extension,
		MatchOn:     ".asc",
		Description: "Potential cryptographic key bundle",
		Comment:     "",
	},
	SimpleFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     "otr.private_key",
		Description: "Pidgin OTR private key",
		Comment:     "",
	},
	SimpleFileSignature{
		Part:        fileSignatureTypes.Extension,
		MatchOn:     ".ovpn",
		Description: "OpenVPN client configuration file",
		Comment:     "",
	},
	SimpleFileSignature{
		Part:        fileSignatureTypes.Extension,
		MatchOn:     ".cscfg",
		Description: "Azure service configuration schema file",
		Comment:     "",
	},
	SimpleFileSignature{
		Part:        fileSignatureTypes.Extension,
		MatchOn:     ".rdp",
		Description: "Remote Desktop connection file",
		Comment:     "",
	},
	SimpleFileSignature{
		Part:        fileSignatureTypes.Extension,
		MatchOn:     ".mdf",
		Description: "Microsoft SQL database file",
		Comment:     "",
	},
	SimpleFileSignature{
		Part:        fileSignatureTypes.Extension,
		MatchOn:     ".sdf",
		Description: "Microsoft SQL server compact database file",
		Comment:     "",
	},
	SimpleFileSignature{
		Part:        fileSignatureTypes.Extension,
		MatchOn:     ".sqlite",
		Description: "SQLite database file",
		Comment:     "",
	},
	SimpleFileSignature{
		Part:        fileSignatureTypes.Extension,
		MatchOn:     ".bek",
		Description: "Microsoft BitLocker recovery key file",
		Comment:     "",
	},
	SimpleFileSignature{
		Part:        fileSignatureTypes.Extension,
		MatchOn:     ".tpm",
		Description: "Microsoft BitLocker Trusted Platform Module password file",
		Comment:     "",
	},
	SimpleFileSignature{
		Part:        fileSignatureTypes.Extension,
		MatchOn:     ".fve",
		Description: "Windows BitLocker full volume encrypted data file",
		Comment:     "",
	},
	SimpleFileSignature{
		Part:        fileSignatureTypes.Extension,
		MatchOn:     ".jks",
		Description: "Java keystore file",
		Comment:     "",
	},
	SimpleFileSignature{
		Part:        fileSignatureTypes.Extension,
		MatchOn:     ".psafe3",
		Description: "Password Safe database file",
		Comment:     "",
	},
	SimpleFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     "secret_token.rb",
		Description: "Ruby On Rails secret token configuration file",
		Comment:     "If the Rails secret token is known, it can allow for remote code execution (http://www.exploit-db.com/exploits/27527/)",
	},
	SimpleFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     "carrierwave.rb",
		Description: "Carrierwave configuration file",
		Comment:     "Can contain credentials for cloud storage systems such as Amazon S3 and Google Storage",
	},
	SimpleFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     "database.yml",
		Description: "Potential Ruby On Rails database configuration file",
		Comment:     "Can contain database credentials",
	},
	SimpleFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     "omniauth.rb",
		Description: "OmniAuth configuration file",
		Comment:     "The OmniAuth configuration file can contain client application secrets",
	},
	SimpleFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     "settings.py",
		Description: "Django configuration file",
		Comment:     "Can contain database credentials, cloud storage system credentials, and other secrets",
	},
	SimpleFileSignature{
		Part:        fileSignatureTypes.Extension,
		MatchOn:     ".agilekeychain",
		Description: "1Password password manager database file",
		Comment:     "Feed it to Hashcat and see if you're lucky",
	},
	SimpleFileSignature{
		Part:        fileSignatureTypes.Extension,
		MatchOn:     ".keychain",
		Description: "Apple Keychain database file",
		Comment:     "",
	},
	SimpleFileSignature{
		Part:        fileSignatureTypes.Extension,
		MatchOn:     ".pcap",
		Description: "Network traffic capture file",
		Comment:     "",
	},
	SimpleFileSignature{
		Part:        fileSignatureTypes.Extension,
		MatchOn:     ".gnucash",
		Description: "GnuCash database file",
		Comment:     "",
	},
	SimpleFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     "jenkins.plugins.publish_over_ssh.BapSshPublisherPlugin.xml",
		Description: "Jenkins publish over SSH plugin file",
		Comment:     "",
	},
	SimpleFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     "credentials.xml",
		Description: "Potential Jenkins credentials file",
		Comment:     "",
	},
	SimpleFileSignature{
		Part:        fileSignatureTypes.Extension,
		MatchOn:     ".kwallet",
		Description: "KDE Wallet Manager database file",
		Comment:     "",
	},
	SimpleFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     "LocalSettings.php",
		Description: "Potential MediaWiki configuration file",
		Comment:     "",
	},
	SimpleFileSignature{
		Part:        fileSignatureTypes.Extension,
		MatchOn:     ".tblk",
		Description: "Tunnelblick VPN configuration file",
		Comment:     "",
	},
	SimpleFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     "Favorites.plist",
		Description: "Sequel Pro MySQL database manager bookmark file",
		Comment:     "",
	},
	SimpleFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     "configuration.user.xpl",
		Description: "Little Snitch firewall configuration file",
		Comment:     "Contains traffic rules for applications",
	},
	SimpleFileSignature{
		Part:        fileSignatureTypes.Extension,
		MatchOn:     ".dayone",
		Description: "Day One journal file",
		Comment:     "Now it's getting creepy...",
	},
	SimpleFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     "journal.txt",
		Description: "Potential jrnl journal file",
		Comment:     "Now it's getting creepy...",
	},
	SimpleFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     "knife.rb",
		Description: "Chef Knife configuration file",
		Comment:     "Can contain references to Chef servers",
	},
	SimpleFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     "proftpdpasswd",
		Description: "cPanel backup ProFTPd credentials file",
		Comment:     "Contains usernames and password hashes for FTP accounts",
	},
	SimpleFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     "robomongo.json",
		Description: "Robomongo MongoDB manager configuration file",
		Comment:     "Can contain credentials for MongoDB databases",
	},
	SimpleFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     "filezilla.xml",
		Description: "FileZilla FTP configuration file",
		Comment:     "Can contain credentials for FTP servers",
	},
	SimpleFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     "recentservers.xml",
		Description: "FileZilla FTP recent servers file",
		Comment:     "Can contain credentials for FTP servers",
	},
	SimpleFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     "ventrilo_srv.ini",
		Description: "Ventrilo server configuration file",
		Comment:     "Can contain passwords",
	},
	SimpleFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     "terraform.tfvars",
		Description: "Terraform variable config file",
		Comment:     "Can contain credentials for terraform providers",
	},
	SimpleFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     ".exports",
		Description: "Shell configuration file",
		Comment:     "Shell configuration files can contain passwords, API keys, hostnames and other goodies",
	},
	SimpleFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     ".functions",
		Description: "Shell configuration file",
		Comment:     "Shell configuration files can contain passwords, API keys, hostnames and other goodies",
	},
	SimpleFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     ".extra",
		Description: "Shell configuration file",
		Comment:     "Shell configuration files can contain passwords, API keys, hostnames and other goodies",
	},
	PatternFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     regexp.MustCompile(`^.*_rsa$`),
		Description: "Private SSH key",
		Comment:     "",
	},
	PatternFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     regexp.MustCompile(`^.*_dsa$`),
		Description: "Private SSH key",
		Comment:     "",
	},
	PatternFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     regexp.MustCompile(`^.*_ed25519$`),
		Description: "Private SSH key",
		Comment:     "",
	},
	PatternFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     regexp.MustCompile(`^.*_ecdsa$`),
		Description: "Private SSH key",
		Comment:     "",
	},
	PatternFileSignature{
		Part:        fileSignatureTypes.Path,
		MatchOn:     regexp.MustCompile(`\.?ssh/config$`),
		Description: "SSH configuration file",
		Comment:     "",
	},
	PatternFileSignature{
		Part:        fileSignatureTypes.Extension,
		MatchOn:     regexp.MustCompile(`^key(pair)?$`),
		Description: "Potential cryptographic private key",
		Comment:     "",
	},
	PatternFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     regexp.MustCompile(`^\.?(bash_|zsh_|sh_|z)?history$`),
		Description: "Shell command history file",
		Comment:     "",
	},
	PatternFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     regexp.MustCompile(`^\.?mysql_history$`),
		Description: "MySQL client command history file",
		Comment:     "",
	},
	PatternFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     regexp.MustCompile(`^\.?psql_history$`),
		Description: "PostgreSQL client command history file",
		Comment:     "",
	},
	PatternFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     regexp.MustCompile(`^\.?pgpass$`),
		Description: "PostgreSQL password file",
		Comment:     "",
	},
	PatternFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     regexp.MustCompile(`^\.?irb_history$`),
		Description: "Ruby IRB console history file",
		Comment:     "",
	},
	PatternFileSignature{
		Part:        fileSignatureTypes.Path,
		MatchOn:     regexp.MustCompile(`\.?purple/accounts\.xml$`),
		Description: "Pidgin chat client account configuration file",
		Comment:     "",
	},
	PatternFileSignature{
		Part:        fileSignatureTypes.Path,
		MatchOn:     regexp.MustCompile(`\.?xchat2?/servlist_?\.conf$`),
		Description: "Hexchat/XChat IRC client server list configuration file",
		Comment:     "",
	},
	PatternFileSignature{
		Part:        fileSignatureTypes.Path,
		MatchOn:     regexp.MustCompile(`\.?irssi/config$`),
		Description: "Irssi IRC client configuration file",
		Comment:     "",
	},
	PatternFileSignature{
		Part:        fileSignatureTypes.Path,
		MatchOn:     regexp.MustCompile(`\.?recon-ng/keys\.db$`),
		Description: "Recon-ng web reconnaissance framework API key database",
		Comment:     "",
	},
	PatternFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     regexp.MustCompile(`^\.?dbeaver-data-sources.xml$`),
		Description: "DBeaver SQL database manager configuration file",
		Comment:     "",
	},
	PatternFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     regexp.MustCompile(`^\.?muttrc$`),
		Description: "Mutt e-mail client configuration file",
		Comment:     "",
	},
	PatternFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     regexp.MustCompile(`^\.?s3cfg$`),
		Description: "S3cmd configuration file",
		Comment:     "",
	},
	PatternFileSignature{
		Part:        fileSignatureTypes.Path,
		MatchOn:     regexp.MustCompile(`\.?aws/credentials$`),
		Description: "AWS CLI credentials file",
		Comment:     "",
	},
	PatternFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     regexp.MustCompile(`^sftp-config(\.json)?$`),
		Description: "SFTP connection configuration file",
		Comment:     "",
	},
	PatternFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     regexp.MustCompile(`^\.?trc$`),
		Description: "T command-line Twitter client configuration file",
		Comment:     "",
	},
	PatternFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     regexp.MustCompile(`^\.?gitrobrc$`),
		Description: "Well, this is awkward... Gitrob configuration file",
		Comment:     "",
	},
	PatternFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     regexp.MustCompile(`^\.?(bash|zsh|csh)rc$`),
		Description: "Shell configuration file",
		Comment:     "Shell configuration files can contain passwords, API keys, hostnames and other goodies",
	},
	PatternFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     regexp.MustCompile(`^\.?(bash_|zsh_)?profile$`),
		Description: "Shell profile configuration file",
		Comment:     "Shell configuration files can contain passwords, API keys, hostnames and other goodies",
	},
	PatternFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     regexp.MustCompile(`^\.?(bash_|zsh_)?aliases$`),
		Description: "Shell command alias configuration file",
		Comment:     "Shell configuration files can contain passwords, API keys, hostnames and other goodies",
	},
	PatternFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     regexp.MustCompile(`config(\.inc)?\.php$`),
		Description: "PHP configuration file",
		Comment:     "",
	},
	PatternFileSignature{
		Part:        fileSignatureTypes.Extension,
		MatchOn:     regexp.MustCompile(`^key(store|ring)$`),
		Description: "GNOME Keyring database file",
		Comment:     "",
	},
	PatternFileSignature{
		Part:        fileSignatureTypes.Extension,
		MatchOn:     regexp.MustCompile(`^kdbx?$`),
		Description: "KeePass password manager database file",
		Comment:     "Feed it to Hashcat and see if you're lucky",
	},
	PatternFileSignature{
		Part:        fileSignatureTypes.Extension,
		MatchOn:     regexp.MustCompile(`^sql(dump)?$`),
		Description: "SQL dump file",
		Comment:     "",
	},
	PatternFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     regexp.MustCompile(`^\.?htpasswd$`),
		Description: "Apache htpasswd file",
		Comment:     "",
	},
	PatternFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     regexp.MustCompile(`^(\.|_)?netrc$`),
		Description: "Configuration file for auto-login process",
		Comment:     "Can contain username and password",
	},
	PatternFileSignature{
		Part:        fileSignatureTypes.Path,
		MatchOn:     regexp.MustCompile(`\.?gem/credentials$`),
		Description: "Rubygems credentials file",
		Comment:     "Can contain API key for a rubygems.org account",
	},
	PatternFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     regexp.MustCompile(`^\.?tugboat$`),
		Description: "Tugboat DigitalOcean management tool configuration",
		Comment:     "",
	},
	PatternFileSignature{
		Part:        fileSignatureTypes.Path,
		MatchOn:     regexp.MustCompile(`doctl/config.yaml$`),
		Description: "DigitalOcean doctl command-line client configuration file",
		Comment:     "Contains DigitalOcean API key and other information",
	},
	PatternFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     regexp.MustCompile(`^\.?git-credentials$`),
		Description: "git-credential-store helper credentials file",
		Comment:     "",
	},
	PatternFileSignature{
		Part:        fileSignatureTypes.Path,
		MatchOn:     regexp.MustCompile(`config/hub$`),
		Description: "GitHub Hub command-line client configuration file",
		Comment:     "Can contain GitHub API access token",
	},
	PatternFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     regexp.MustCompile(`^\.?gitconfig$`),
		Description: "Git configuration file",
		Comment:     "",
	},
	PatternFileSignature{
		Part:        fileSignatureTypes.Path,
		MatchOn:     regexp.MustCompile(`\.?chef/(.*)\.pem$`),
		Description: "Chef private key",
		Comment:     "Can be used to authenticate against Chef servers",
	},
	PatternFileSignature{
		Part:        fileSignatureTypes.Path,
		MatchOn:     regexp.MustCompile(`etc/shadow$`),
		Description: "Potential Linux shadow file",
		Comment:     "Contains hashed passwords for system users",
	},
	PatternFileSignature{
		Part:        fileSignatureTypes.Path,
		MatchOn:     regexp.MustCompile(`etc/passwd$`),
		Description: "Potential Linux passwd file",
		Comment:     "Contains system user information",
	},
	PatternFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     regexp.MustCompile(`^\.?dockercfg$`),
		Description: "Docker configuration file",
		Comment:     "Can contain credentials for public or private Docker registries",
	},
	PatternFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     regexp.MustCompile(`^\.?npmrc$`),
		Description: "NPM configuration file",
		Comment:     "Can contain credentials for NPM registries",
	},
	PatternFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     regexp.MustCompile(`^\.?env$`),
		Description: "Environment configuration file",
		Comment:     "",
	},
	PatternFileSignature{
		Part:        fileSignatureTypes.Path,
		MatchOn:     regexp.MustCompile(`credential`),
		Description: "Contains word: credential",
		Comment:     "",
	},
	PatternFileSignature{
		Part:        fileSignatureTypes.Path,
		MatchOn:     regexp.MustCompile(`password`),
		Description: "Contains word: password",
		Comment:     "",
	},
	SimpleFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     "credentials.json",
		Description: "Google Cloud Platform service account credentials keyfile",
		Comment:     "GCP service account credentials can be activated using the gcloud command from Google's Cloud SDK (https://cloud.google.com/sdk/gcloud/reference/auth/activate-service-account)",
	},
	PatternFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     regexp.MustCompile(`^.*-[a-f0-9]{12}\.json$`),
		Description: "Google Cloud Platform service account credentials keyfile",
		Comment:     "GCP service account credentials can be activated using the gcloud command from Google's Cloud SDK (https://cloud.google.com/sdk/gcloud/reference/auth/activate-service-account)",
	},
	SimpleFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     "adc.json",
		Description: "Legacy Google Cloud Platform service account credentials keyfile",
		Comment:     "GCP service account credentials can be activated using the gcloud command from Google's Cloud SDK (https://cloud.google.com/sdk/gcloud/reference/auth/activate-service-account)",
	},
	SimpleFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     "credentials.db",
		Description: "Google Cloud Platform gcloud credential database",
		Comment:     "sqlite database containing credentials used by the gcloud command from Google's Cloud SDK",
	},
	SimpleFileSignature{
		Part:        fileSignatureTypes.Filename,
		MatchOn:     ".boto",
		Description: "Legacy Google Cloud Platform gcloud credential database",
		Comment:     "File containing credentials used by the gcloud command from Google's Cloud SDK",
	},
}