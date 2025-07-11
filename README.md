# go-fiber-pg
go-fiber server is created along postgress as a database for API call handling

Action  	     Method	   URL	                                                                        Body Example (if needed)  ||
HomeCheck	     GET	     http://localhost:3000/	                                                        none  ||
Seed Records	 GET	     http://localhost:3000/seed	                                                    none  ||
Get All	       GET	     http://localhost:3000/records                                                	none    ||
Create	       POST	     http://localhost:3000/record	                                                  { "key": "key_1001", "value": "some value" }  ||
Get One	       GET	     http://localhost:3000/record/key_1001                                         	 none   ||
Update	       PUT	     http://localhost:3000/record/key_1001	                                         { "value": "updated value" }  ||
Delete	       DELETE  	 http://localhost:3000/record/key_1001	                                          none    ||
