package main

func main() {
	jack := NewUser("Jack", "Solomon", "sol.jack@proton.me")
	jack.displayUser()

	will := NewAdmin("Hawkins", "HNL", "Will", "Byers", "byers.001Will@proton.me")
	will.displayAdmin()
	will.user.displayUser()
}
