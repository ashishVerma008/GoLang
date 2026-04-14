package main

import "github.com/gofiber/fiber/v3"

type Data struct{
	Message string `json:"message"`
	Status int	   `json:"status"`
}

type Query struct{
	Name string `json:"name"`
	Age int		`json:"age"`
}

func main() {
	app := fiber.New()

    // app.Get("/",func(c fiber.Ctx) error{
	// 	data:=Data{
	// 		Message: "All are perfect",
	// 		Status: 200,
	// 	}
    //     // return c.SendString("Hello, World!")
	// 	return c.JSON(data )
	// })


// get params
//    app.Get("/:id",func(c fiber.Ctx) error{
// 	userid:=c.Params("id")
// 	return c.JSON(fiber.Map{
// 		"message":"We get user id",
// 		 "userid":userid,
// 	})
//    })


// Get query
// app.Get("/",func(c fiber.Ctx)error{
// 	name:=c.Query("name")
// 	data:=Data{
// 		Message: name,
// 		Status: 200,
// 	}
// 	return c.JSON(data)
// })

//Multiple queries
// app.Get("/",func(c fiber.Ctx) error{
// 	query  := c.Queries()
// 	return c.JSON(query)
// }) 


// Body
app.Post("/",func(c fiber.Ctx) error{
	data:=new(Data)
	err:=c.Bind().Body(data)
	if err!=nil {
		return err
	}
	return c.JSON(data)
})
	app.Listen(":3000")

}