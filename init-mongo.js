db.createUser(
    {
        user: "root",
        pwd: "point1234",
        roles: [
            {
                role: "readWrite",
                db: "todopoint"
            }
        ]
    }
);