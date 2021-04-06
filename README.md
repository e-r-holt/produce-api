# produce-api
API for interacting with produce database

## API Methods
### GET 

<details>
    <summary> / </summary>

Returns all records in the produce database
</details>
<details>
    <summary> /{produce-code} </summary>
Returns a specific record identified by the given produce code
</details>

### POST

<details>
    <summary> / </summary>
Returns the record(s) that were added

Payload: List of JSON objects defining produce. Keys: Name, Produce Code, Unit Price
<pre>
    <code>
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
</details>