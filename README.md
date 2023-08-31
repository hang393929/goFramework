# goFramework
配置：config.yaml
app:
  env: 'local'
  port: 8888
  app_name: 'gin-app'
  app_url: 'http://localhost'

log:
  level: 'info'
  root_dir: './storage/logs'
  filename: 'app.log'
  format: ''
  show_line: true
  max_backups: 3
  max_size: 500
  max_age: 28
  compress: true

database:
  driver: 'mysql'
  host: ''
  port: 63824
  database: 'xzq_test'
  username: 'root'
  password: ''
  charset: 'utf8mb4'
  max_idle_conns: 10
  max_open_conns: 100
  log_mode: 'info'
  enable_file_log_writer: true
  log_filename: 'sql.log'

jwt:
  secret: '3Bde3BGEbYqtqyEUzW3ry8jKFcaPH17fRmTmqE7MDr05Lwj95uruRKrrkb44TJ4s'
  jwt_ttl: 43200
  jwt_blacklist_grace_period: 10
  refresh_grace_period: 1800

redis:
  host: ''
  port: 23000
  db: 0
  password: ''

storage:
  default: 'local'
  disks:
    local:
      root_dir: './storage/app'
      app_url: 'http://localhost:8888/storage'
    tx_oss:
      access_key_id: 'access_key_id'
      access_key_secret: 'access_key_secret'
      bucket: 'bucket'
      endpoint: 'endpoint'
      is_ssl: true
      is_private: false
    qi_niu:
      access_key: 'access_key'
      bucket: 'bucket'
      domain: 'domain'
      secret_key: 'secret_key'
      is_ssl: true
      is_private: false
