# Animals interface API

## **Endpoints Documentation**

### **1. Process Animal**

**Endpoint**: `/process`  
**Method**: `GET`

**Description**: Adds a new animal to the in-memory database.

#### **Request**

- `curl http:localhost:8080/process?type=dog&name=milo&attr="5-inches"`

- **Response**:
  ```json
  {
    "isAnimalType": true
  }
  ```

### **2. Describe Animal**

**Endpoint**: `/describe`  
**Method**: `GET`

**Description**: Adds a new animal to the in-memory database.

#### **Request**

- `curl http:localhost:8080/describe?type=cat&name=kity&attr="gray"`

- **Response**:
  ```json
  {
    "Type": "Cat"
  }
  ```

### **3. Get Animal by Id**

**Endpoint**: `/get?id={id}`  
**Method**: `GET`

**Description**: gets an animal by its Id from the in-memory database.

#### **Request**

- `curl http:localhost:8080/get?id=kitty`

- **Response**:
  ```json
  {
    "name": "kitty",
    "Color": "gray",
    "Methods": {
      "Speak": "Meow!",
      "Category": "Mammal"
    }
  }
  ```

  ### **4. Delete Animal by Id**

**Endpoint**: `/delete?id={id}`  
**Method**: `GET`

**Description**: Remove an animal to the in-memory database.

#### **Request**

- `curl http:localhost:8080/delete?id=kitty`

- **Response**:
  ```json
  {
    "Sucess": "Animal with given id kitty deleted"
  }
  ```


    ### **5. get Animald**

**Endpoint**: `/animals?type={type}`  
**Method**: `GET`

**Description**: get all animals in Db, if the type is pased as param filter by types

#### **Request**

- `curl http:localhost:8080/animals`

- **Response**:
  ```json
  {
      "data": [
    {
      "name": "kitty",
      "Color": "gray",
      "Methods": {
        "Speak": "Meow!",
        "Category": "Mammal"
      }
    },
    {
      "name": "milo",
      "Breed": "5- inches",
      "Methods": {
        "Speak": "Wooof!",
        "Category": "Mammal"
      }
    },
    {
      "name": "paco",
      "WingSpan": "56",
      "Methods": {
        "Speak": "Chirp!",
        "Category": "Aves"
      }
    }
  ]
  }
  ```


 
