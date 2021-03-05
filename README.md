Usage:

POST http://localhost:8080/ 

    {
        "x": 10,
        "y": 2,
        "z": 4
    }
        
returns

    {
        "state_changes": [
            {
                "name": "Fill Y",
                "state": {
                    "x": 0,
                    "y": 2
                }
            },
            {
                "name": "Transfer Y to X",
                "state": {
                    "x": 2,
                    "y": 0
                }
            },
            {
                "name": "Fill Y",
                "state": {
                    "x": 2,
                    "y": 2
                }
            },
            {
                "name": "Transfer Y to X",
                "state": {
                    "x": 4,
                    "y": 0
                }
            }
        ],
        "length": 4
    }
    
for input 

    {
        "x": 6,
        "y": 3,
        "z": 1
    }
    
returns 

    No Solutions