package request_strings

var Register string = `mutation Register(
	$first_name: String!, 
	$last_name: String, 
	$username: String, 
	$email: String!, 
	$password: String) {
  insert_users_one(object: {first_name: $first_name, last_name: $last_name, username: $username, email: $email, password: $password}) {
    id
    first_name
    last_name
    username
    email
  }
}`
var SearchUserByUsernameOrEmail string = `query SearchUser($Identifier: String!) {
  users(where: {_or: [ {email: {_eq: $Identifier}}, {username: {_eq: $Identifier}}]}) {
    id
    first_name
    last_name
    username
    email
    password
  }}
`
var SearchUserByEmail string = `query MyQuery($email: String) {
  users(where: {email: {_eq: $email}}) {
    id
    first_name
    last_name
    username
    email
  }
}
`

var SearchUserByUsername string = `query MyQuery($username: String) {
  users(where: {username: {_eq: $username}}) {
    id
    first_name
    last_name
    username
    email
  }
}
`
