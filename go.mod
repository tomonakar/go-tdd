module github.com/tomonakar/go-tdd

go 1.15

replace go-tdd/integers => ./integers

replace go-tdd/iteration => ./for

replace go-tdd/arrays => ./arrays

replace go-tdd/structs => ./structs

replace go-tdd/pointers => ./pointers

replace go-tdd/maps => ./maps

replace go-tdd/di => ./di

replace go-tdd/concurrency => ./concurrency

replace go-tdd/race => ./race

replace go-tdd/reflection => ./reflection

require (
	github.com/kisielk/errcheck v1.4.0 // indirect
	golang.org/x/mod v0.4.0 // indirect
	golang.org/x/tools v0.0.0-20201211185031-d93e913c1a58 // indirect
)
