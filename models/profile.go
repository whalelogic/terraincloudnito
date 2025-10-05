// Package models creates structures to hold data or dynamically render data.
package models


type Experience struct {
        CompanyName string
        Role        string
        StartDate   string
        EndDate     string
        Description string
}

type Skill struct {
        Name        string
        Proficiency string
}

type Language struct {
        Name        string
        Proficiency string
}

type Interest struct {
        Name string
}

type Education struct {
        SchoolName   string
        Degree       string
        FieldOfStudy string
}

type Profile struct {
        Name        string
        Title       string
        Bio         string
        AvatarURL   string
        Educations  []Education
        Experiences []Experience
        Skills      []Skill
        Languages   []Language
        Interests   []Interest
}


var MyProfile = Profile{
        Name:      "Keith Thomson",
        Title:     "Software and Cloud Solutions Developer",
        Bio:       "Passionate developer with experience in Go, Information Systems design, and Application development.",
        AvatarURL: "https://gravatar.com/avatar/24289225bb8abcf8cdd1fe4d5d20db4d?s=400&d=robohash&r=x",
        Educations: []Education{
                {
                        SchoolName:   "Southern Connecticut State University",
                        Degree:       "Bachelor of Science",
                        FieldOfStudy: "Computer Science - Information Systems",
                },
                {
                        SchoolName:   "Connecticut State College",
                        Degree:       "Associate of Science",
                        FieldOfStudy: "Cloud Computing & Virtualization",
                },
                {
                        SchoolName:   "Coursera",
                        Degree:       "Professional Certificate",
                        FieldOfStudy: "Google Data Analytics",
                },
                {
                        SchoolName:   "Connecticut State College",
                        Degree:       "Industry Certification",
                        FieldOfStudy: "Amazon Web Services Cloud Practioner",
                },
                {
                        SchoolName:   "W3 Schools",
                        Degree:       "Professional Certificate",
                        FieldOfStudy: "C++ Programming Language",
                },
                {
                        SchoolName:   "Python Institute",
                        Degree:       "PCEP ",
                        FieldOfStudy: "Programming Languages",
                },
        },
        Experiences: []Experience{
                {
                        CompanyName: "WhalerAPI",
                        Role:        "Founder & Lead Developer",
                        StartDate:   "Mar 2025",
                        EndDate:     "Present",
                        Description: "Designing Cloud Solutions & Developing scalable backend services using Go and Python.",
                },
                {
                        CompanyName: "PennyBorne Development LLC",
                        Role:        "Product QA Lead - Content Management System",
                        StartDate:   "Mar 2022",
                        EndDate:     "Present",
                        Description: "Developing scalable backend services using Go.",
                },
                {
                        CompanyName: "The Spice and Tea Exchange",
                        Role:        "Supervisor - Assistant Manager",
                        StartDate:   "May 2016",
                        EndDate:     "Oct 2020",
                        Description: "Coordinating events and promotions to generate sales revenue.",
               	}, 
				{
                        CompanyName: "The Magetti Group LLC",
                        Role:        "Accounting Clerk",
                        StartDate:   "Feb 2012",
                        EndDate:     "Feb 2016",
                        Description: "Developing scalable backend services using Go.",
                },
                {
                        CompanyName: "US Army",
                        Role:        "11B Infantry",
                        StartDate:   "Sep 2007",
                        EndDate:     "Sep 2011",
                        Description: "Squad Leader - Deployed with 3rd Heavy Brigade Combat Team 🪖, 1st Cavalry Division - Operation Iraqi Freedom November 2008 - January 2010 🎖️🎖️Mosul, Ninevah Province.",
                },
        },
        Skills: []Skill{
                {Name: "Go Programming Language", Proficiency: "Professional"},
                {Name: "Python ", Proficiency: "Advanced"},
                {Name: "Bash", Proficiency: "Expert"},
                {Name: "AWS - Azure - GCP", Proficiency: "Professional"},
                {Name: "Linux 🐧Operating Systems", Proficiency: "Professional"},
                {Name: "Docker", Proficiency: "Advanced"},
                {Name: "Kubernetes", Proficiency: "Intermediate"},
                {Name: "Terraform", Proficiency: "Advanced"},
                {Name: "Data Engineering", Proficiency: "Advanced"},
                {Name: "Application Development", Proficiency: "Advanced"},
                {Name: "Microservices", Proficiency: "Professional"},
                {Name: "Cloud Architecture", Proficiency: "Expert"},
                {Name: "RESTful API", Proficiency: "Expert"},
                {Name: "Software Design Patterns", Proficiency: "Advanced"},
                {Name: "IoT", Proficiency: "Advanced"},
                {Name: "Neural Networks and Machine Learning", Proficiency: "Advanced"},
				
        },
        Languages: []Language{
                {Name: "English", Proficiency: "Native"},
                {Name: "Spanish", Proficiency: "Conversational"},
                {Name: "German", Proficiency: "Basic"},
        },
        Interests: []Interest{
                {Name: "Open Source Contribution"},
                {Name: "Cloud Computing"},
                {Name: "Family time 🐒"},
                {Name: "Traveling 🛫"},
        },
}


