package db

import (
  "gorm.io/gorm"
)

type Profile struct {
  gorm.Model
  UserType           string         `json:"usertype" gorm:"size:30"`
  Username           string         `json:"username" gorm:"size:20"`
  Password           string         `json:"pxwd" gorm:"size:150"`
}

type Pupil struct {
  gorm.Model
  Profile            Profile        `json:"profile"`
  profileID          uint           `json:"profileid"`
  Firstname          string         `json:"firstname" gorm:"size:20"`
  Lastname           string         `json:"lastname" gorm:"size:20"`
  Phone              string         `json:"phone" gorm:"size:20"`
  UserCode           string         `json:"usercode" gorm:"size:20"`
  Deactivated        bool           `json:"deactivated" gorm:"default:false"`
}

type Teacher struct {
  gorm.Model
  Profile            Profile        `json:"profile"`
  profileID          uint           `json:"profileid"`
  Firstname          string         `json:"firstname" gorm:"size:20"`
  Lastname           string         `json:"lastname" gorm:"size:20"`
}

/*
        
        $assignmentList_table_definition = "
        id int auto_increment primary key not null,
        expired boolean default false,
        startDate date,
        endDate date,
        startTime time,
        endTime time,
        submitedAt timestamp
        ";
        
        $assignment_table_definition = "
        id int auto_increment primary key not null,
        assignmentListID int not null,
        characterToDraw  char not null,
        expectedSolution varchar(50) not null
        createdAt timestamp
        ";
        
        $activationRequest_table_definition = "
        id int auto_increment primary key not null,
        pupilID int not null,
        requestMessage varchar(200) default '',
        submitedAt timestamp
        ";
        
        $solutions_table_definition = "
        id int auto_increment primary key not null,
        pupilID int not null,
        assignmentID int not null,
        assignmentListID int not null,
        attemptTimeTakenInSeconds int,
        characterToDraw char not null,
        answer varchar(50) not null,
        autoAttachedScore int,
        submitedAt timestamp
        ";

*/
 
  





