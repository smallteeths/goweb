package service

import (
	"github.com/gin-gonic/gin"
	. "tool-backend/handler"
    "tool-backend/pkg/error"
    "tool-backend/model"
    "fmt"
    "github.com/satori/go.uuid"
    "io"
    "bytes"
    "io/ioutil"
    "os"
    "log"
    "github.com/spf13/viper"
    "text/template"
    "bufio"
    "strings"
    "encoding/json"
    "os/exec"
    "net/http"
    "github.com/gorilla/websocket"
    "encoding/gob"
)
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type  LoginLinkData struct {
	Link 	string `json:"link" form:"link"`
}

func UploadHandler(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
	   SendResponse(c, errno.ErrBind, nil)
	   return
	}
	fileName := header.Filename
	fmt.Println(file, err, fileName)
	//创建文件
	u1 := uuid.Must(uuid.NewV4())
	fmt.Printf("UUIDv4: %s\n", u1)
  
	newFileName := "logo" + u1.String() + ".svg"
	out, err := os.Create("static/uploadfile/" + newFileName)
	//注意此处的 static/uploadfile/ 不是/static/uploadfile/
	if err != nil {
        log.Fatal(err)
        SendResponse(c, errno.ErrBind, nil)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
        log.Fatal(err)
        SendResponse(c, errno.ErrBind, nil)
    }
    
    u := model.TemplateVariable{
		FileName: newFileName,
    }

	SendResponse(c, nil, u)
}

func UploadBackgroundHandler(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
	   SendResponse(c, errno.ErrBind, nil)
	   return
	}
	fileName := header.Filename
	fmt.Println(file, err, fileName)
	//创建文件
	u1 := uuid.Must(uuid.NewV4())
	fmt.Printf("UUIDv4: %s\n", u1)
  
	newFileName := "login" + u1.String() + ".svg"
	out, err := os.Create("static/uploadfile/" + newFileName)
	//注意此处的 static/uploadfile/ 不是/static/uploadfile/
	if err != nil {
        log.Fatal(err)
        SendResponse(c, errno.ErrBind, nil)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
        log.Fatal(err)
        SendResponse(c, errno.ErrBind, nil)
    }
    
    u := model.TemplateVariable{
		LoginBgFileName: newFileName,
    }

	SendResponse(c, nil, u)
}

func Save(c *gin.Context) {
    var r model.TemplateVariable
    
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
    }

	u := model.TemplateVariable{
        FileName: r.FileName,
        LoginBgFileName: r.LoginBgFileName,
        LinkData: r.LinkData,
        VariablesData: r.VariablesData,
        LoginrecordData: r.LoginrecordData,
    }

    if r.LinkData != "" {
        ChangeFooterFile(c, r.LinkData)
        info, err := ioutil.ReadFile("static/rancherfile/footer.hbs")

        if err != nil {
            fmt.Println(err)
            SendResponse(c, errno.ErrBind, nil)
            return
        }
        
        out := []byte(info)
        ioutil.WriteFile(viper.GetString("rancherfooteraddr"), out, 0655)
    }

    if r.VariablesData != "" {
        ChangeTheme(c, r.VariablesData)
    }

    if r.LoginrecordData != "" {
        ChangeLoginRecord(c, r.LoginrecordData)
    }

    if r.FileName != "" {
        info, err := ioutil.ReadFile("static/uploadfile/" + r.FileName)

        if err != nil {
            fmt.Println(err)
            SendResponse(c, errno.ErrBind, nil)
            return
        }
        
        out := []byte(info)
        ioutil.WriteFile(viper.GetString("logoaddr"), out, 0655)
    }

    if r.LoginBgFileName != "" {
        info, err := ioutil.ReadFile("static/uploadfile/" + r.LoginBgFileName)

        if err != nil {
            fmt.Println(err)
            SendResponse(c, errno.ErrBind, nil)
            return
        }
        
        out := []byte(info)
        ioutil.WriteFile(viper.GetString("loginbgaddr"), out, 0655)
    }

    file, err := os.Create(viper.GetString("savefileaddr"))
    if err != nil {
        fmt.Println(err)
    }

    enc := gob.NewEncoder(file)
    err2 := enc.Encode(u)

    if err2 != nil {
        SendResponse(c, errno.ErrBind, nil)
		return
    }
    
    SendResponse(c, nil, nil)
}
 
func ChangeFooterFile(c *gin.Context, list string) {

    fmt.Printf("string: %s\n", list)
    
    var linkDataList []model.LinkData

    json.Unmarshal([]byte(list), &linkDataList)

    var concats string

    for _,item:= range linkDataList {
        data := model.LinkData{
            LinkName: item.LinkName,
            LinkAddr: item.LinkAddr,
        }

        s := `<a style="color: #fff" role="button" class="btn btn-sm bg-transparent" target="blank" rel="noreferrer noopener" href="{{.LinkAddr}}">{{.LinkName}}</a>`

        t, err := template.New("test").Parse(s)

        if err != nil {
            fmt.Println("Fatal error ", err.Error())
            SendResponse(c, errno.ErrBind, nil)
            return
        }
    
        buf := new(bytes.Buffer)
        err = t.Execute(buf, data)
    
        if err != nil {
            fmt.Println("Fatal error ", err.Error())
            SendResponse(c, errno.ErrBind, nil)
            return
        }
    
        concats += buf.String()
    }

    fmt.Printf(concats)

	in, err := os.Open(viper.GetString("footerfileaddr"))
	if err != nil {
		fmt.Println("open file fail:", err)
		os.Exit(-1)
	}
	defer in.Close()
 
    del := os.Remove(viper.GetString("footerfileaddr")+".hbs");
    if del != nil {
        fmt.Println(del);
    }

	out, err := os.OpenFile(viper.GetString("footerfileaddr")+".hbs", os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		fmt.Println("Open write file fail:", err)
		os.Exit(-1)
	}
	defer out.Close()
 
	br := bufio.NewReader(in)

	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read err:", err)
			os.Exit(-1)
		}
		newLine := strings.Replace(string(line), "rancher-tool-wsy-link", concats, -1)
		_, err = out.WriteString(newLine + "\n")
		if err != nil {
			fmt.Println("write to file fail:", err)
			os.Exit(-1)
		}
	}
	fmt.Println("FINISH!")
}

func ChangeTheme(c *gin.Context, themeString string) {

    var themeColor model.ThemeColor

    json.Unmarshal([]byte(themeString), &themeColor)

    fmt.Printf("string: %s\n", themeColor.Primary)

    info, err := ioutil.ReadFile(viper.GetString("themefileaddr"))

    if err != nil {
        fmt.Println(err)
        SendResponse(c, errno.ErrBind, nil)
        return
    }
    
    out := []byte(info)

    t, err := template.New("test").Parse(string(out))

    if err != nil {
        fmt.Println("Fatal error ", err.Error())
        SendResponse(c, errno.ErrBind, nil)
        return
    }

    buf := new(bytes.Buffer)

    err = t.Execute(buf, themeColor)
    
    if err != nil {
        fmt.Println("Fatal error ", err.Error())
        SendResponse(c, errno.ErrBind, nil)
        return
    }

    name := viper.GetString("rancherthemeaddr")

    ioutil.WriteFile(name, []byte(buf.String()), 0655)

}

func ChangeLoginRecord(c *gin.Context, LoginrecordString string) {

    var loginrecordData model.LoginrecordData

    json.Unmarshal([]byte(LoginrecordString), &loginrecordData)

    s := `<div style="position: fixed; padding: 0px 20px; bottom: 20px; left: 0px; width: 100%; text-align: right; background: rgba(204,204,204, .3);">
        <a style="color: #3497DA" role="button" class="btn btn-sm bg-transparent" target="blank" rel="noreferrer noopener" href="{{.LinkAddr}}">{{.LinkName}}</a>
    </div>`
    
    t, err := template.New("test").Parse(s)

    if err != nil {
        fmt.Println("Fatal error ", err.Error())
        SendResponse(c, errno.ErrBind, nil)
        return
    }

    buf := new(bytes.Buffer)
    err = t.Execute(buf, loginrecordData)

    if err != nil {
        fmt.Println("Fatal error ", err.Error())
        SendResponse(c, errno.ErrBind, nil)
        return
    }

    fmt.Println("Link FINISH!")
    
    info, err := ioutil.ReadFile(viper.GetString("loginfileaddr"))

    if err != nil {
        fmt.Println(err)
        SendResponse(c, errno.ErrBind, nil)
        return
    }
    
    out := []byte(info)

    link := LoginLinkData{
        Link: buf.String(),
    }

    tRead, err := template.New("test").Delims("[[", "]]").Parse(string(out))

    if err != nil {
        fmt.Println("Fatal error ", err.Error())
        SendResponse(c, errno.ErrBind, nil)
        return
    }

    bufRead := new(bytes.Buffer)

    err = tRead.Execute(bufRead, link)

    if err != nil {
        fmt.Println("Fatal error ", err.Error())
        SendResponse(c, errno.ErrBind, nil)
        return
    }

    name := viper.GetString("rancherloginfileaddr")

    ioutil.WriteFile(name, []byte(bufRead.String()), 0655)

}

func SelectTemplateVariable(c *gin.Context)  {
    var u model.TemplateVariable

    file, err := os.Open(viper.GetString("savefileaddr"))

    if err != nil {
        SendResponse(c, errno.ErrBind, nil)
        fmt.Println(err)
    }
    dec := gob.NewDecoder(file)
    err2 := dec.Decode(&u)
    
    if err2 != nil {
        SendResponse(c, errno.ErrBind, nil)
        return
    }

    SendResponse(c, nil, u)
}

func Test(c *gin.Context) {

    ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)

    if err != nil {
		SendResponse(c, errno.ErrBind, nil)
    }
    
    defer ws.Close()

    for {
		//读取ws中的数据
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
        }
        
        fmt.Printf("string: %s\n", string(message))

		if string(message) == "build" {
            fmt.Printf("string: %s\n", "build start")
    
            command := viper.GetString("buildfile")

            cmd := exec.Command("/bin/sh", command)
        
            stdout, _ := cmd.StdoutPipe()
            stderr, _ := cmd.StderrPipe()
        
            if err := cmd.Start(); err != nil {
                log.Printf("Error starting command: %s......", err.Error())
            }
        
            asyncLog(stdout, mt, ws)
            asyncLog(stderr, mt, ws)
        
            if err := cmd.Wait(); err != nil {
                log.Printf("Error waiting for command execution: %s......", err.Error())
                ws.WriteMessage(mt, []byte("Failed build"))
                break
            }
            
            ws.WriteMessage(mt, []byte("Done build"))
		}
	}

}

func asyncLog(reader io.ReadCloser,mt int, ws *websocket.Conn) error {
	cache := ""
	buf := make([]byte, 1024)
	for {
		num, err := reader.Read(buf)
		if err != nil && err!=io.EOF{
			return err
        }
        if num == 0 {
            break
        }
		if num > 0 {
			b := buf[:num]
			s := strings.Split(string(b), "\n")
			line := strings.Join(s[:len(s)-1], "\n") 
            fmt.Printf("%s%s\n", cache, line)
            err = ws.WriteMessage(mt, []byte(line))
			cache = s[len(s)-1]
		}
    }

	return nil
}
