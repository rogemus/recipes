package main

type sessionKey = string

const userIdSessionKey = sessionKey("authenticatedUserID")
const userNameSessionKey = sessionKey("authenticatedUserName")
