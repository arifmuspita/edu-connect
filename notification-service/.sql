Table notifications {
  notification_id integer [primary key]
  email varchar(100) 
  subject varchar(100)
  message text
  status varchar(10) [default: "pending"]
  created_at timestamp
}