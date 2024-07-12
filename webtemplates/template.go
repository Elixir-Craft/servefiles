package webtemplates

var Home = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Home</title>
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

        .card-container {
            display: flex;
            flex-wrap: wrap;
            gap: 20px;
        }

        .card {
            background-color: #fff;
            border: 1px solid #ddd;
            border-radius: 8px;
            box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
            padding: 20px;
            width: calc(50% - 20px);
            box-sizing: border-box;
            transition: background-color 0.3s;
            display: flex;
            align-items: center;
            max-width: 300px;
        }

        .card:hover {
            background-color: #f1f1f1;
        }

        .icon {
            display: inline-block;
            width: 20px;
            height: 20px;
            margin-right: 10px;
            background-size: contain;
            background-repeat: no-repeat;
        }

        .directory-icon {
            background-image: url('data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiIHN0YW5kYWxvbmU9Im5vIj8+Cgo8IS0tIExpY2Vuc2U6IFBELiBNYWRlIGJ5IHBhb21lZGlhOiBodHRwczovL2dpdGh1Yi5jb20vcGFvbWVkaWEvc21hbGwtbi1mbGF0IC0tPgo8c3ZnIHdpZHRoPSI4MDBweCIgaGVpZ2h0PSI4MDBweCIgdmlld0JveD0iMCAwIDI0IDI0IiB4bWxuczpyZGY9Imh0dHA6Ly93d3cudzMub3JnLzE5OTkvMDIvMjItcmRmLXN5bnRheC1ucyMiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyIgdmVyc2lvbj0iMS4xIiB4bWxuczpjYz0iaHR0cDovL2NyZWF0aXZlY29tbW9ucy5vcmcvbnMjIiB4bWxuczpkYz0iaHR0cDovL3B1cmwub3JnL2RjL2VsZW1lbnRzLzEuMS8iPgogPGcgdHJhbnNmb3JtPSJ0cmFuc2xhdGUoMCAtMTAyOC40KSI+CiAgPHBhdGggZD0ibTIgMTAzMy40Yy0xLjEwNDYgMC0yIDAuOS0yIDJ2MTRjMCAxLjEgMC44OTU0MyAyIDIgMmgyMGMxLjEwNSAwIDItMC45IDItMnYtMTRjMC0xLjEtMC44OTUtMi0yLTJoLTIweiIgZmlsbD0iIzI5ODBiOSIvPgogIDxwYXRoIGQ9Im0zIDEwMjkuNGMtMS4xMDQ2IDAtMiAwLjktMiAydjE0YzAgMS4xIDAuODk1NCAyIDIgMmgxMSA1IDJjMS4xMDUgMCAyLTAuOSAyLTJ2LTktM2MwLTEuMS0wLjg5NS0yLTItMmgtMi01LTFsLTMtMmgtN3oiIGZpbGw9IiMyOTgwYjkiLz4KICA8cGF0aCBkPSJtMjMgMTA0Mi40di04YzAtMS4xLTAuODk1LTItMi0yaC0xMS01LTJjLTEuMTA0NiAwLTIgMC45LTIgMnY4aDIyeiIgZmlsbD0iI2JkYzNjNyIvPgogIDxwYXRoIGQ9Im0yIDEwMzMuNGMtMS4xMDQ2IDAtMiAwLjktMiAydjYgMSA2YzAgMS4xIDAuODk1NDMgMiAyIDJoMjBjMS4xMDUgMCAyLTAuOSAyLTJ2LTYtMS02YzAtMS4xLTAuODk1LTItMi0yaC0yMHoiIGZpbGw9IiMzNDk4ZGIiLz4KIDwvZz4KPC9zdmc+');
        }

        .github-icon {
            background-image: url('data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0idXRmLTgiPz4KPCEtLSBMaWNlbnNlOiBBcGFjaGUuIE1hZGUgYnkgSWNvbnNjb3V0OiBodHRwczovL2dpdGh1Yi5jb20vSWNvbnNjb3V0L3VuaWNvbnMgLS0+CjxzdmcgZmlsbD0iIzAwMDAwMCIgd2lkdGg9IjgwMHB4IiBoZWlnaHQ9IjgwMHB4IiB2aWV3Qm94PSIwIDAgMjQgMjQiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyIgZGF0YS1uYW1lPSJMYXllciAxIj48cGF0aCBkPSJNMTIsMi4yNDY3QTEwLjAwMDQyLDEwLjAwMDQyLDAsMCwwLDguODM3NTIsMjEuNzM0MTljLjUuMDg3NTIuNjg3NS0uMjEyNDcuNjg3NS0uNDc1LDAtLjIzNzQ5LS4wMTI1MS0xLjAyNS0uMDEyNTEtMS44NjI0OUM3LDE5Ljg1OTE5LDYuMzUsMTguNzg0MjMsNi4xNSwxOC4yMjE3M0EzLjYzNiwzLjYzNiwwLDAsMCw1LjEyNSwxNi44MDkyYy0uMzUtLjE4NzUtLjg1LS42NS0uMDEyNTEtLjY2MjQ4QTIuMDAxMTcsMi4wMDExNywwLDAsMSw2LjY1LDE3LjE3MTY5YTIuMTM3NDIsMi4xMzc0MiwwLDAsMCwyLjkxMjQ4LjgyNUEyLjEwMzc2LDIuMTAzNzYsMCwwLDEsMTAuMiwxNi42NTkyM2MtMi4yMjUtLjI1LTQuNTUtMS4xMTI1NC00LjU1LTQuOTM3NWEzLjg5MTg3LDMuODkxODcsMCwwLDEsMS4wMjUtMi42ODc1LDMuNTkzNzMsMy41OTM3MywwLDAsMSwuMS0yLjY1cy44Mzc0Ny0uMjYyNTEsMi43NSwxLjAyNWE5LjQyNzQ3LDkuNDI3NDcsMCwwLDEsNSwwYzEuOTEyNDgtMS4zLDIuNzUtMS4wMjUsMi43NS0xLjAyNWEzLjU5MzIzLDMuNTkzMjMsMCwwLDEsLjEsMi42NSwzLjg2OSwzLjg2OSwwLDAsMSwxLjAyNSwyLjY4NzVjMCwzLjgzNzQ3LTIuMzM3NTIsNC42ODc1LTQuNTYyNSw0LjkzNzVhMi4zNjgxNCwyLjM2ODE0LDAsMCwxLC42NzUsMS44NWMwLDEuMzM3NTItLjAxMjUxLDIuNDEyNDgtLjAxMjUxLDIuNzUsMCwuMjYyNTEuMTg3NS41NzUuNjg3NS40NzVBMTAuMDA1MywxMC4wMDUzLDAsMCwwLDEyLDIuMjQ2N1oiLz48L3N2Zz4=');
        }

        .card a {
            text-decoration: none;
            color: #007BFF;
            font-size: 18px;
        }

        .card a:hover {
            text-decoration: underline;
        }

    </style>
</head>
<body>
    <h1>ServeIt is running</h1>
    <div class="card-container">
        <div class="card">
            <span class="icon directory-icon"></span>
            <a href="/files">Files</a>
        </div>
        <div class="card">
            <span class="icon github-icon"></span>
            <a href="https://github.com/Elixir-Craft/serveIt" target="_blank">GitHub Repo</a>
        </div>
    </div>
</body>
</html>
`

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
            display: inline-block;
            width: 20px;
            height: 20px;
            margin-right: 10px;
            background-size: contain;
            background-repeat: no-repeat;
        }

        .file {
            color: #6c757d;
        }
        .file-icon {
            background-image: url('data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0idXRmLTgiPz4KPCEtLSBMaWNlbnNlOiBNSVQuIE1hZGUgYnkgbWljaGFlbGFtcHI6IGh0dHBzOi8vZ2l0aHViLmNvbS9taWNoYWVsYW1wci9qYW0gLS0+CjxzdmcgZmlsbD0iIzAwMDAwMCIgd2lkdGg9IjgwMHB4IiBoZWlnaHQ9IjgwMHB4IiB2aWV3Qm94PSItNCAtMiAyNCAyNCIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIiBwcmVzZXJ2ZUFzcGVjdFJhdGlvPSJ4TWluWU1pbiIgY2xhc3M9ImphbSBqYW0tZmlsZSI+PHBhdGggZD0nTTEwLjI5OCAySDNhMSAxIDAgMCAwLTEgMXYxNGExIDEgMCAwIDAgMSAxaDEwYTEgMSAwIDAgMCAxLTFWNC45NjFMMTAuMjk4IDJ6TTMgMGg4bDUgNHYxM2EzIDMgMCAwIDEtMyAzSDNhMyAzIDAgMCAxLTMtM1YzYTMgMyAwIDAgMSAzLTN6Jy8+PC9zdmc+');
        }
        .directory {
            color: #5c67f2;
        }
        .directory-icon {
            background-image: url('data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiIHN0YW5kYWxvbmU9Im5vIj8+Cgo8IS0tIExpY2Vuc2U6IFBELiBNYWRlIGJ5IHBhb21lZGlhOiBodHRwczovL2dpdGh1Yi5jb20vcGFvbWVkaWEvc21hbGwtbi1mbGF0IC0tPgo8c3ZnIHdpZHRoPSI4MDBweCIgaGVpZ2h0PSI4MDBweCIgdmlld0JveD0iMCAwIDI0IDI0IiB4bWxuczpyZGY9Imh0dHA6Ly93d3cudzMub3JnLzE5OTkvMDIvMjItcmRmLXN5bnRheC1ucyMiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyIgdmVyc2lvbj0iMS4xIiB4bWxuczpjYz0iaHR0cDovL2NyZWF0aXZlY29tbW9ucy5vcmcvbnMjIiB4bWxuczpkYz0iaHR0cDovL3B1cmwub3JnL2RjL2VsZW1lbnRzLzEuMS8iPgogPGcgdHJhbnNmb3JtPSJ0cmFuc2xhdGUoMCAtMTAyOC40KSI+CiAgPHBhdGggZD0ibTIgMTAzMy40Yy0xLjEwNDYgMC0yIDAuOS0yIDJ2MTRjMCAxLjEgMC44OTU0MyAyIDIgMmgyMGMxLjEwNSAwIDItMC45IDItMnYtMTRjMC0xLjEtMC44OTUtMi0yLTJoLTIweiIgZmlsbD0iIzI5ODBiOSIvPgogIDxwYXRoIGQ9Im0zIDEwMjkuNGMtMS4xMDQ2IDAtMiAwLjktMiAydjE0YzAgMS4xIDAuODk1NCAyIDIgMmgxMSA1IDJjMS4xMDUgMCAyLTAuOSAyLTJ2LTktM2MwLTEuMS0wLjg5NS0yLTItMmgtMi01LTFsLTMtMmgtN3oiIGZpbGw9IiMyOTgwYjkiLz4KICA8cGF0aCBkPSJtMjMgMTA0Mi40di04YzAtMS4xLTAuODk1LTItMi0yaC0xMS01LTJjLTEuMTA0NiAwLTIgMC45LTIgMnY4aDIyeiIgZmlsbD0iI2JkYzNjNyIvPgogIDxwYXRoIGQ9Im0yIDEwMzMuNGMtMS4xMDQ2IDAtMiAwLjktMiAydjYgMSA2YzAgMS4xIDAuODk1NDMgMiAyIDJoMjBjMS4xMDUgMCAyLTAuOSAyLTJ2LTYtMS02YzAtMS4xLTAuODk1LTItMi0yaC0yMHoiIGZpbGw9IiMzNDk4ZGIiLz4KIDwvZz4KPC9zdmc+');
        }

        .file-size {
            margin-left: auto;
            color: #6c757d;
            font-size: 14px;
        }
    </style>
</head>
<body>
<h1>ServeIt</h1>
    <h2>Directory Listing</h2>
    <ul>

        <!-- .. -->
        <li>
        <span class="icon directory-icon"></span>
        <a href="../">../</a>
        </li>



        {{range .Files}}
        <li>
            {{if .IsDir}}
            <span class="icon directory-icon"></span>
                <a href="{{$.CurrentPath}}/{{.Name}}">{{.Name}}/</a>
            {{else}}
            <span class="icon file-icon"></span>
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
