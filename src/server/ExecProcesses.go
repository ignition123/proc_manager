package server

import(
	"lib"
    "os/exec"
    "fmt"
    "objects"
    "strings"
    "time"
    "strconv"
    "encoding/json"
    "ByteBuffer"
    "syscall"
)

func ExecProcesses(){

	defer lib.Handlepanic()

	// iterating over the config process list

	for index := range objects.Conf.App{

		// splitting the arguments with spaces

		arguments := strings.Split(*objects.Conf.App[index].Cmd, " ")

		// checking if time exists in config

		if objects.Conf.App[index].Sleep != nil{
			time.Sleep(time.Duration(*objects.Conf.App[index].Sleep) * time.Second)
		}

		// spawnning the processess

		go Exec(arguments[0], arguments[0:], objects.Conf.App[index])
	}

	for{
		time.Sleep(1000000 * time.Hour)
	}
}

func Exec(app string, args []string, conf *objects.Applications){

	defer lib.Handlepanic()

	RESTART:

	// creating the command for execution

	cmd := exec.Command(app, args...)

	// creating a std pipeline

	stdout, err := cmd.StdoutPipe()

	if err != nil{
		fmt.Println(string(objects.ColorRed), err)
		time.Sleep(5 * time.Second)
		goto RESTART
	}

	// starting the process

	cmd.SysProcAttr = &syscall.SysProcAttr{
	    Pdeathsig: syscall.SIGKILL,
	}

	cmd.Start()

	fmt.Println(string(objects.ColorGreen), "Running application: ", *conf.Name)

	// checking if schedular exists

	go runSchedular(conf, cmd)

	out := make([]byte, 128)

	// listening the process

	for{

		n, err := stdout.Read(out)

		if err != nil{
			fmt.Println(string(objects.ColorRed), err)
			break
		}

		// printing the logs

		fmt.Println(string(objects.ColorBlue), string(out[:n]))

		if TCPCon != nil{

			func(){

				defer lib.Handlepanic()

				msgStruct := &objects.Msg{
					Msg: string(out[:n]),
					PsName: *conf.Name,
					Ip: serverIP,
					LUT: time.Now().String(),
					PsPath: *conf.Path,
				}

				msg, err := json.Marshal(msgStruct)

				if err != nil{
					fmt.Println(string(objects.ColorRed), err)
					return
				}

				bb := &ByteBuffer.Buffer{
					Endian:"big",
				}

				bb.PutShort(len(msg))
				bb.Put(msg)

				_, tcpErr := TCPCon.Write(bb.Array())

				if tcpErr != nil{
					fmt.Println(string(objects.ColorRed), tcpErr)
					return
				}

			}()

		}

	}

	// application is crashed restarting the process

	fmt.Println(string(objects.ColorRed), "Process crashed: ", *conf.Name)

	cmd.Wait()

	time.Sleep(5 * time.Second)

	fmt.Println(string(objects.ColorGreen), "Restarting process: ", *conf.Name)

	goto RESTART
}

// running schedular

func runSchedular(conf *objects.Applications, cmd *exec.Cmd){

	defer lib.Handlepanic()

	// spliting time

	times := strings.Split(*conf.Restart, ":")

	// converting string to int
	hour, err := strconv.Atoi(times[0])

	if err != nil{
		fmt.Println(string(objects.ColorRed), err)
		return
	}

	// converting string to int
	min, err := strconv.Atoi(times[1])

	if err != nil{
		fmt.Println(string(objects.ColorRed), err)
		return
	}

	// checking schedular every minutes
	for{

		// getting current hour and min of clock
		hours, minutes, _ := time.Now().Clock()

		if hours == hour && minutes == min{

			cmd.Process.Kill()

			break
		}

		time.Sleep(1 * time.Minute)

		// if tcp connection is not null
		if TCPCon != nil{

			func(){

				// writing status of the process to server
				defer lib.Handlepanic()

				statStruct := &objects.Stat{
					PsName: *conf.Name,
					Ip: serverIP,
					LUT: time.Now().String(),
					PsPath: *conf.Path,
				}

				msg, err := json.Marshal(statStruct)

				if err != nil{
					fmt.Println(string(objects.ColorRed), err)
					return
				}


				bb := &ByteBuffer.Buffer{
					Endian:"big",
				}

				bb.PutShort(len(msg))
				bb.Put(msg)

				_, tcpErr := TCPCon.Write(bb.Array())

				if tcpErr != nil{
					fmt.Println(string(objects.ColorRed), tcpErr)
					return
				}

			}()

		}
	}
}