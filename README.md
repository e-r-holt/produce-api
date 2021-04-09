# produce-api
API for interacting with produce database


## Start hosting!
Image is available on Docker hub
```
docker pull deemaflotchy/produce-api
docker run -d -p 3000:3000 deemaflotchy/produce-api
```

## API Methods
### GET 

<details>
    <summary> / </summary>

Returns all records in the produce database in JSON list format

<pre>
    <code>
//Response content
[
    {
        "Name" : "Dragonfruit",
        "Produce Code" : "1234-1234-1234-1234",
        "Unit Price" : 3.14
    },
    {
        "Name" : "Orange",
        "Produce Code" : "4312-4312-4312-4321",
        "Unit Price" : 2.14
    }
]
    </code>
</pre>
</details>
</details>
<details>
    <summary> /{produce-code} </summary>
Returns a specific record identified by the given produce code

<pre>
    <code>
//Response content
[
    {
        "Name" : "Dragonfruit",
        "Produce Code" : "1234-1234-1234-1234",
        "Unit Price" : 3.14
    }
]
    </code>
</pre>
</details>
</details>

### POST

<details>
    <summary> / </summary>
Returns the record(s) that were added

<pre>
    <code>
//Payload
[
    {
        "Name" : "Dragonfruit",
        "Produce Code" : "1234-1234-1234-1234",
        "Unit Price" : 3.14
    }
]
    </code>
</pre>
</details>

### DELETE
<details>
    <summary> /{produce-code} </summary>
Deletes a specific record identified by the given produce code

Returns the record that was deleted
<pre>
    <code>
//Response content
[
    {
        "Name" : "Dragonfruit",
        "Produce Code" : "1234-1234-1234-1234",
        "Unit Price" : 3.14
    }
]
    </code>
</pre>
</details>
</details>