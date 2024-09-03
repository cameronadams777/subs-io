package controllers

import (
	"app/views/components"

	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/labstack/echo/v4"
)

type SubtitlesController struct{}

const MAX_UPLOAD_SIZE = 1024 * 1024 * 1024

func output_command_logs(cmd *exec.Cmd) {
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		fmt.Println("Error creating StdoutPipe:", err)
		return
	}
	if err = cmd.Start(); err != nil {
		fmt.Println("Error starting command:", err)
		return
	}
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		fmt.Print(string(tmp))
		if err != nil {
			break
		}
	}
}

func generate_subtitles(file_path string) {
	whisper_command := "whisperx " + file_path + ".mp4 --model medium.en --output_dir ./uploads --align_model WAV2VEC2_ASR_LARGE_LV60K_960H --batch_size 4 --compute_type float32"
	whisper_command_arr := strings.Split(whisper_command, " ")
	whisper_exec := exec.Command(whisper_command_arr[0], whisper_command_arr[1:]...)
	output_command_logs(whisper_exec)

	ffmpeg_command := "ffmpeg -i " + file_path + ".mp4 -vf subtitles=" + file_path + ".srt " + file_path + "_subbed.mp4"
	ffmpeg_command_arr := strings.Split(ffmpeg_command, " ")
	ffmpeg_exec := exec.Command(ffmpeg_command_arr[0], ffmpeg_command_arr[1:]...)
	output_command_logs(ffmpeg_exec)
}

func generate_subbed_video(file_path string) {
}

func (sc *SubtitlesController) HandleSubtitlesCreate(c echo.Context) error {
	c.Request().Body = http.MaxBytesReader(c.Response(), c.Request().Body, MAX_UPLOAD_SIZE)
	if err := c.Request().ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		return render(c, components.FlashMessage(components.FlashMessageProps{
			Message: err.Error(),
		}))
	}

	file, err := c.FormFile("video")
	if err != nil {
		return render(c, components.FlashMessage(components.FlashMessageProps{
			Message: err.Error(),
		}))
	}

	src, err := file.Open()
	if err != nil {
		return render(c, components.FlashMessage(components.FlashMessageProps{
			Message: err.Error(),
		}))
	}
	defer src.Close()

	file_name := strings.ToLower(strings.ReplaceAll(file.Filename, " ", "_"))

	// Destination
	dst, err := os.Create("uploads/" + file_name)
	if err != nil {
		return render(c, components.FlashMessage(components.FlashMessageProps{
			Message: err.Error(),
		}))
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return render(c, components.FlashMessage(components.FlashMessageProps{
			Message: err.Error(),
		}))
	}

	path := "uploads/" + strings.Split(file_name, ".")[0]

	generate_subtitles(path)

	c.Response().Header().Set("HX-Retarget", "#upload-form")
	return c.String(http.StatusOK, "Success")
}
