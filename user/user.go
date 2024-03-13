package user

type User struct {
	Username      string   `yaml:"username"`
	FullName      string   `yaml:"full_name"`
	PublicKeyPem  string   `yaml:"public_key_pem"`
	PrivateKeyPem string   `yaml:"private_key_pem"`
	Followers     []string `yaml:"followers"`
}
