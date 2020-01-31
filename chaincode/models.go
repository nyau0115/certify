package main

type Organisation struct { //Network Administrator
    Id string `json:"id"`
    Name string `json:"name"`
    Region string `json:"region"`
}

type User struct {
    Id string `json:"id"`
    Name string `json:"name"`
    Gender string  `json:"gender"`
    NationalId string `json:"nationalId"`
    PhoneNumber string `json:"phoneNumber"`
    Address string `json:"address"`
    Organisation Organisation `json:"organisation"` // Organisation Administrator Immutable
    Title string `json:"address"`   //
    UserType string `json:"userType"` // 1: Student, 2: Instructor, 3: Endorser
}
