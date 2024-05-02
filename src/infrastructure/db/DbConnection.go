package db

import "fmt"

type PostgresConnection struct {
}

func (p *PostgresConnection) Connect() {
	fmt.Println("Conectando ao banco de dados PostgreSQL...")
	// Implemente aqui a lógica de conexão com o PostgreSQL
}

type MySqlConnection struct {
}

func (m *MySqlConnection) Connect() {
	fmt.Println("Conectando ao banco de dados MySQL...")
	// Implemente aqui a lógica de conexão com o MySQL
}

type OracleConnection struct {
}

func (o *OracleConnection) Connect() {
	fmt.Println("Conectando ao banco de dados Oracle...")
}

type S3Connection struct {
}

func (s *S3Connection) Connect() {
	fmt.Println("Conectando ao S3...")
}

type DbConnection struct {
	Conn interface{}
}

func NewDbConnection(connectionType string) *DbConnection {
	switch connectionType {
	case "postgres":
		return &DbConnection{Conn: &PostgresConnection{}}
	case "mysql":
		return &DbConnection{Conn: &MySqlConnection{}}
	case "oracle":
		return &DbConnection{Conn: &OracleConnection{}}
	case "s3":
		return &DbConnection{Conn: &S3Connection{}}
	default:
		panic("invalid type")
	}
}

func NewDbConnectionBase() *DbConnection {
	return &DbConnection{}
}

func (d *DbConnection) Connect() {
	switch conn := d.Conn.(type) {
	case *PostgresConnection:
		conn.Connect()
	case *MySqlConnection:
		conn.Connect()
	case *OracleConnection:
		conn.Connect()
	case *S3Connection:
		conn.Connect()
	default:
		panic("Invalid type")
	}
}
