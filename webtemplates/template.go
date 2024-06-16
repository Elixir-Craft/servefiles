package webtemplates

var Index = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Directory Listing</title>
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0-beta3/css/all.min.css">
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f4f4f4;
            color: #333;
            margin: 0;
            padding: 20px;
        }
        h1 {
            font-size: 24px;
            margin-bottom: 20px;
        }
        ul {
            list-style-type: none;
            padding: 0;
        }
        li {
            display: flex;
            align-items: center;
            padding: 10px;
            background-color: #fff;
            border: 1px solid #ddd;
            margin-bottom: 10px;
            border-radius: 4px;
            transition: background-color 0.3s;
        }
        li:hover {
            background-color: #f1f1f1;
        }
        a {
            text-decoration: none;
            color: #007BFF;
            flex-grow: 1;
        }
        a:hover {
            text-decoration: underline;
        }
        .icon {
            font-size: 18px;
            width: 20px;
            margin-right: 10px;
        }
        .file {
            color: #6c757d;
        }
        .directory {
            color: #5c67f2;
        }
        .file-size {
            margin-left: auto;
            color: #6c757d;
            font-size: 14px;
        }
    </style>
</head>
<body>
    <h1>Directory Listing</h1>
    <ul>

        <!-- .. -->
        <li>
            <i class="fas fa-folder icon directory"></i>
            <a href="{{.ParentPath}}">../</a>
        </li>



        {{range .Files}}
        <li>
            {{if .IsDir}}
                <i class="fas fa-folder icon directory"></i>
                <a href="{{$.CurrentPath}}/{{.Name}}">{{.Name}}/</a>
            {{else}}
                <i class="fas fa-file icon file"></i>
                <a href="{{$.CurrentPath}}/{{.Name}}">{{.Name}}</a>
                {{if .Size}}
                <span class="file-size">{{.Size}}</span>
                {{end}}
            {{end}}
        </li>
        {{else}}
        <li><em>No files or directories found.</em></li>
        {{end}}
    </ul>
</body>
</html>`

var Auth = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Password Required</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f4f4f4;
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            margin: 0;
        }
        form {
            background-color: white;
            padding: 20px;
            border-radius: 8px;
            box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
        }
        h1 {
            color: #333;
            text-align: center;
        }
        label {
            margin-bottom: 10px;
            color: #666;
            display: block;
        }
        input[type="password"] {
            width: calc(100% - 22px);
            padding: 10px;
            margin-top: 5px;
            margin-bottom: 20px;
            border: 1px solid #ddd;
            border-radius: 4px;
            box-sizing: border-box; /* Adds padding and border to the element's total width and height */
        }
        input[type="submit"] {
            width: 100%;
            padding: 10px;
            border: none;
            border-radius: 4px;
            color: white;
            background-color: #5c67f2;
            cursor: pointer;
        }
        input[type="submit"]:hover {
            background-color: #4a54e1;
        }
    </style>
</head>
<body>
    <form action="" method="post">
        <h1>Password Required</h1>
        <label for="password">Password:</label>
        <input type="password" id="password" name="password">
        <input type="submit" value="Submit">
    </form>
</body>
</html>
`
