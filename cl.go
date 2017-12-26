package main 

import(
  
	"fmt"   
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"math/rand"
	"time"
	"sort"
)

type Team struct{
	
		Id  bson.ObjectId `json:"_id" bson:"_id"` 
		Name string `json:"Name" bson:"Name"` 
		Pot int `json:"Pot" bson:"Pot"`
		Country int `json:"Country" bson:"Country"`
		GroupPoints int
		Average int
}


type Group struct {
	
		Id  bson.ObjectId `json:"_id" bson:"_id"` 
		Teams []Team `json:"Teams" bson:"Teams"` 
		
	
}


var Msession *mgo.Session

func main(){

	Msession = BuildMongo()   
	Draw()
	VS()
	VSAfterParty()
	VSLast4()
	Final()
	
	 
}


var GroupA Group
var GroupB Group
var GroupC Group
var GroupD Group 

var AfterParty []Team
var Last4 []Team
var FinalMatch []Team


func Draw(){

	var pot1Pool []Team
	var pot2Pool []Team
	var pot3Pool []Team
	var pot4Pool []Team


	

	pot1Pool = GetAllTeamsByPot(1)
	pot2Pool = GetAllTeamsByPot(2)
	pot3Pool = GetAllTeamsByPot(3)
	pot4Pool = GetAllTeamsByPot(4)

	 
	numbersPot1 := []int{0,1,2,3}
	numbersPot2 := []int{0,1,2,3}
	numbersPot3 := []int{0,1,2,3}
	numbersPot4 := []int{0,1,2,3}
	n := 3


	for i:=0;i<=3;i++{

		
		var randTeam1 int
		if i!=3 {
			randTeam1 = randInt(0, n-i)
		}else{
			randTeam1 = 0 
		}

		rand.Seed(time.Now().UnixNano())

		var randTeam2 int
		if i!=3 {
			randTeam2= randInt(0, n-i)
		}else{
			randTeam2 = 0 
		}

		rand.Seed(time.Now().UnixNano())
		var randTeam3 int
		if i!=3 {
			randTeam3 = randInt(0, n-i)
		}else{
			randTeam3 = 0 
		}

		rand.Seed(time.Now().UnixNano())

		var randTeam4 int
		if i!=3 {
			randTeam4 = randInt(0, n-i)
		}else{
			randTeam4 = 0 
		}


		if i==0{

			GroupA.Teams = append(GroupA.Teams,pot1Pool[numbersPot1[randTeam1]])
			GroupA.Teams = append(GroupA.Teams,pot2Pool[numbersPot2[randTeam2]])
			GroupA.Teams = append(GroupA.Teams,pot3Pool[numbersPot3[randTeam3]])
			GroupA.Teams = append(GroupA.Teams,pot4Pool[numbersPot4[randTeam4]])
		}
		if i==1{

			GroupB.Teams = append(GroupB.Teams,pot1Pool[numbersPot1[randTeam1]])
			GroupB.Teams = append(GroupB.Teams,pot2Pool[numbersPot2[randTeam2]])
			GroupB.Teams = append(GroupB.Teams,pot3Pool[numbersPot3[randTeam3]])
			GroupB.Teams = append(GroupB.Teams,pot4Pool[numbersPot4[randTeam4]])
		}
		if i==2{

			GroupC.Teams = append(GroupC.Teams,pot1Pool[numbersPot1[randTeam1]])
			GroupC.Teams = append(GroupC.Teams,pot2Pool[numbersPot2[randTeam2]])
			GroupC.Teams = append(GroupC.Teams,pot3Pool[numbersPot3[randTeam3]])
			GroupC.Teams = append(GroupC.Teams,pot4Pool[numbersPot4[randTeam4]])
		}
		if i==3{

			GroupD.Teams = append(GroupD.Teams,pot1Pool[numbersPot1[randTeam1]])
			GroupD.Teams = append(GroupD.Teams,pot2Pool[numbersPot2[randTeam2]])
			GroupD.Teams = append(GroupD.Teams,pot3Pool[numbersPot3[randTeam3]])
			GroupD.Teams = append(GroupD.Teams,pot4Pool[numbersPot4[randTeam4]])
		}
		
		numbersPot1 = RemoveIndex(numbersPot1,randTeam1)
		numbersPot2 = RemoveIndex(numbersPot2,randTeam2)
		numbersPot3 = RemoveIndex(numbersPot3,randTeam3)
		numbersPot4 = RemoveIndex(numbersPot4,randTeam4) 
	 
	}

	fmt.Println("GROUP-A___________")
	for i:=0;i<=3;i++{
		fmt.Println(GroupA.Teams[i].Name)
	}
	fmt.Println("GROUP-B___________")
	for i:=0;i<=3;i++{
		fmt.Println(GroupB.Teams[i].Name)
	}
	fmt.Println("GROUP-C___________")
	for i:=0;i<=3;i++{
		fmt.Println(GroupC.Teams[i].Name)
	}
	fmt.Println("GROUP-D___________")
	for i:=0;i<=3;i++{
		fmt.Println(GroupD.Teams[i].Name)
	}
 
	fmt.Println("****************___________********************")
	fmt.Println("****************___________********************")
	fmt.Println("****************___________********************")


}

func VS(){

	Groups := []Group{GroupA,GroupB,GroupC,GroupD}

	for _, group := range Groups {
			
		
		for a:=0;a<=3;a++{

			for b:=0;b<=3;b++{
				var teamScore1 int
				var teamScore2 int
				
				teamScore1 = randInt(0,4)
				rand.Seed(time.Now().UnixNano())
				teamScore2 = randInt(0,4)

				if a!=b{
					fmt.Println(group.Teams[a].Name,":",teamScore1," - ",group.Teams[b].Name+":",teamScore2)

					if teamScore1>teamScore2{
						group.Teams[a].GroupPoints = group.Teams[a].GroupPoints +3 
					}else if teamScore2>teamScore1{
						group.Teams[b].GroupPoints = group.Teams[b].GroupPoints +3
					}else{
						group.Teams[a].GroupPoints = group.Teams[a].GroupPoints +1
						group.Teams[b].GroupPoints = group.Teams[b].GroupPoints +1
					}
				}

			}
			
		}

		
		for i:=0;i<=3;i++{
			sort.Slice(group.Teams, func(i, j int) bool {return group.Teams[i].GroupPoints > group.Teams[j].GroupPoints})
			fmt.Println(group.Teams[i].Name , " : ",group.Teams[i].GroupPoints)

			
			 
		}

		AfterParty = append(AfterParty,group.Teams[0])
		AfterParty = append(AfterParty,group.Teams[1])
		fmt.Println("****************___________********************")
		fmt.Println("****************___________********************")
	}
	fmt.Println("AFTER PARTY")
	fmt.Println("****************___________********************")
	fmt.Println("****************___________********************")

	for _, team := range AfterParty{

		fmt.Println(team.Name)

	}
	
 
}

func VSAfterParty(){

	for i:=0;i<=5;i++{
		var teamScore1 int
		var teamScore2_1 int
		var teamScore2_2 int
		
		AfterParty[i].GroupPoints=0
		AfterParty[i+2].GroupPoints=0


		teamScore1 = randInt(0,4)
		rand.Seed(time.Now().UnixNano())
		teamScore2_1 = randInt(0,4)

		if i!=2 && i!=3{
		//MATCH 1  
		fmt.Println(AfterParty[i].Name,":",teamScore1," - ",AfterParty[i+2].Name+":",teamScore2_1)

		if teamScore1>teamScore2_1{
			AfterParty[i].GroupPoints = AfterParty[i].GroupPoints +3 
			AfterParty[i].Average = AfterParty[i].Average +(teamScore1-teamScore2_1) 
			
		}else if teamScore2_1>teamScore1{
			AfterParty[i+2].GroupPoints = AfterParty[i+2].GroupPoints +3
			AfterParty[i+2].Average = AfterParty[i+2].Average +(teamScore2_1-teamScore1) 
			
		}else{
			AfterParty[i].GroupPoints = AfterParty[i].GroupPoints +1
			AfterParty[i+2].GroupPoints = AfterParty[i+2].GroupPoints +1
		}

		//MATCH 2
		teamScore1 = randInt(0,4)
		rand.Seed(time.Now().UnixNano())
		teamScore2_2 = randInt(0,4)

		fmt.Println(AfterParty[i+2].Name,":",teamScore1," - ",AfterParty[i].Name+":",teamScore2_2)

		if teamScore1>teamScore2_2{
			AfterParty[i+2].GroupPoints = AfterParty[i+2].GroupPoints +3 
			AfterParty[i+2].Average = AfterParty[i+2].Average +(teamScore1-teamScore2_2) 
		}else if teamScore2_2>teamScore1{
			AfterParty[i].GroupPoints = AfterParty[i].GroupPoints +3
			AfterParty[i].Average = AfterParty[i].Average +(teamScore2_2-teamScore1) 
		}else{
			AfterParty[i].GroupPoints = AfterParty[i].GroupPoints +1
			AfterParty[i+2].GroupPoints = AfterParty[i+2].GroupPoints +1
		}
		 

		if AfterParty[i].GroupPoints > AfterParty[i+2].GroupPoints{

			Last4 = append(Last4,AfterParty[i])

		}else if AfterParty[i].GroupPoints < AfterParty[i+2].GroupPoints{

			Last4 = append(Last4,AfterParty[i+2])
			
		}else{

			 

			if AfterParty[i].Average>AfterParty[i+2].Average{
				Last4 = append(Last4,AfterParty[i])

			}else if AfterParty[i+2].Average>AfterParty[i].Average{
				Last4 = append(Last4,AfterParty[i+2])
			}else{

				if teamScore2_1>teamScore2_2{
					
									Last4 = append(Last4,AfterParty[i])
								}else{
					
									Last4 = append(Last4,AfterParty[i+2])
									
								}
			}
			 
			


		}

	}


	}

	fmt.Println("LAST 4")
	fmt.Println("****************___________********************")
	fmt.Println("****************___________********************")
	
		for _, team := range Last4{
	
			fmt.Println(team.Name)
	
		}

		

}

func VSLast4(){


	for i:=0;i<=1;i++{
		var teamScore1 int
		var teamScore2_1 int
		var teamScore2_2 int
		
		AfterParty[i].GroupPoints=0
		AfterParty[i+2].GroupPoints=0


		teamScore1 = randInt(0,4)
		rand.Seed(time.Now().UnixNano())
		teamScore2_1 = randInt(0,4)

		 
		//MATCH 1  
		fmt.Println(Last4[i].Name,":",teamScore1," - ",Last4[i+2].Name+":",teamScore2_1)

		if teamScore1>teamScore2_1{
			Last4[i].GroupPoints = Last4[i].GroupPoints +3 
			Last4[i].Average = Last4[i].Average +(teamScore1-teamScore2_1) 
			
		}else if teamScore2_1>teamScore1{
			Last4[i+2].GroupPoints = Last4[i+2].GroupPoints +3
			Last4[i+2].Average = Last4[i+2].Average +(teamScore2_1-teamScore1) 
			
		}else{
			Last4[i].GroupPoints = Last4[i].GroupPoints +1
			Last4[i+2].GroupPoints = Last4[i+2].GroupPoints +1
		}

		//MATCH 2
		teamScore1 = randInt(0,4)
		rand.Seed(time.Now().UnixNano())
		teamScore2_2 = randInt(0,4)

		fmt.Println(Last4[i+2].Name,":",teamScore1," - ",Last4[i].Name+":",teamScore2_2)

		if teamScore1>teamScore2_2{
			Last4[i+2].GroupPoints = Last4[i+2].GroupPoints +3 
			Last4[i+2].Average = Last4[i+2].Average +(teamScore1-teamScore2_2) 
		}else if teamScore2_2>teamScore1{
			Last4[i].GroupPoints = Last4[i].GroupPoints +3
			Last4[i].Average = Last4[i].Average +(teamScore2_2-teamScore1) 
		}else{
			Last4[i].GroupPoints = Last4[i].GroupPoints +1
			Last4[i+2].GroupPoints = Last4[i+2].GroupPoints +1
		}
		 

		if Last4[i].GroupPoints > Last4[i+2].GroupPoints{

			FinalMatch = append(FinalMatch,Last4[i])

		}else if Last4[i].GroupPoints < Last4[i+2].GroupPoints{

			FinalMatch = append(FinalMatch,Last4[i+2])
			
		}else{

			 

			if Last4[i].Average>Last4[i+2].Average{
				FinalMatch = append(FinalMatch,Last4[i])

			}else if Last4[i+2].Average>Last4[i].Average{
				FinalMatch = append(FinalMatch,Last4[i+2])
			}else{

				fmt.Println("teamScore2_1",teamScore2_1)
				fmt.Println("teamScore2_2",teamScore2_2)

				if teamScore2_1>teamScore2_2{
					
					FinalMatch = append(FinalMatch,Last4[i+2])
				}else{
	
					FinalMatch = append(FinalMatch,Last4[i])
					
				}
			}
			 
			


		}

 


	}
	fmt.Println("FINAL")
	fmt.Println("****************___________********************")
	fmt.Println("****************___________********************")
	
		for _, team := range FinalMatch{
	
			fmt.Println(team.Name)
	
		}

	

}


func Final(){
	var teamScore1 int
	var teamScore2 int

	teamScore1 = randInt(0,3)
	rand.Seed(time.Now().UnixNano())
	teamScore2 = randInt(0,3)


	//MATCH 1  
			fmt.Println(FinalMatch[0].Name,":",teamScore1," - ",FinalMatch[1].Name+":",teamScore2)
	
			if teamScore1>teamScore2{
				 
				fmt.Println("****Champion****",FinalMatch[0].Name)
				
			}else if teamScore2>teamScore1{
				 
				fmt.Println("****Champion****",FinalMatch[1].Name)
				 
				
			}else{
				fmt.Println("Extra Time Baby")
				fmt.Println("Penalties")
				
				fmt.Println("****Champion :) ****",FinalMatch[0].Name)
			}



}

func GetAllTeams(){
	
	   c := Msession.DB("ChampionsLeague").C("Teams") 
	   var  teams []Team
	   var err error
	   err = c.Find(nil).All(&teams)
		
   
	   if err != nil {
		   panic(err)
	   }
 
   }



   
   func GetAllTeamsByPot(pot int) []Team{
	
	   c := Msession.DB("ChampionsLeague").C("Teams") 
	   var  teams []Team
	   var err error
	   err = c.Find(bson.M{"Pot": pot}).All(&teams)
		
   
	   if err != nil {
		   panic(err)
	   }
	
	   

	   return teams
   }
   
func randInt(min int, max int) int {
    return min + rand.Intn(max-min)
}

func RemoveIndex(s []int, index int) []int {
    return append(s[:index], s[index+1:]...)
}
func BuildMongo() *mgo.Session{
	
		session, err := mgo.Dial("127.0.0.1")
		if err != nil {
			panic(err)
		}
	
		defer session.Close()
	
		session.SetMode(mgo.Monotonic, true)
	 
		// Collection People
		c := session.DB("ChampionsLeague").C("Teams")
	
		// Index
		index := mgo.Index{
			Key:        []string{"id", "Name","Pot","Country"},
			Unique:     true,
			DropDups:   true,
			Background: true,
			Sparse:     true,
		}
	
		err = c.EnsureIndex(index)
		if err != nil {
			panic(err)
		}
	
		fmt.Println("Build Mongo Db Done!")
		
		return session.Clone()
	
}
	
	 