package longpoll

import (
	"os"
	"strconv"
	"testing"

	"github.com/SevereCloud/vksdk/api"
)

func TestLongpoll_Shutdown(t *testing.T) {
	t.Parallel()

	t.Run("Shutdown", func(t *testing.T) {
		lp := &Longpoll{}
		lp.Shutdown()
		if lp.inShutdown != 1 {
			t.Error("inShutdown != 1")
		}
	})
}

func TestLongpoll_Handler(t *testing.T) {
	t.Parallel()
	// nolint:gocyclo
	lp := &Longpoll{}

	t.Run("FullResponse", func(t *testing.T) {
		lp.FullResponse(func(resp Response) {})
		if len(lp.funcFullResponseList) != 1 {
			t.Error("Want len = 1")
		}
	})
}

func TestNewLongpoll(t *testing.T) {
	t.Parallel()

	groupToken := os.Getenv("GROUP_TOKEN")
	if groupToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	vk := api.NewVK(groupToken)
	badVk := api.NewVK("")
	groupID, _ := strconv.Atoi(os.Getenv("GROUP_ID"))

	type args struct {
		vk      *api.VK
		groupID int
	}

	tests := []struct {
		name string
		args args
		// wantLp  Longpoll
		wantErr bool
	}{
		{
			name: "Init error",
			args: args{
				vk:      badVk,
				groupID: 0,
			},
			wantErr: true,
		},
		{
			name: "Init good",
			args: args{
				vk:      vk,
				groupID: groupID,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			_, err := NewLongpoll(tt.args.vk, tt.args.groupID)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewLongpoll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(gotLp, tt.wantLp) {
			// 	t.Errorf("NewLongpoll() = %v, want %v", gotLp, tt.wantLp)
			// }
		})
	}
}

func TestLongpoll_checkResponse(t *testing.T) {
	groupToken := os.Getenv("GROUP_TOKEN")
	if groupToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	vk := api.NewVK(groupToken)
	groupID, _ := strconv.Atoi(os.Getenv("GROUP_ID"))
	lp, _ := NewLongpoll(vk, groupID)

	tests := []struct {
		name        string
		argResponse Response
		wantErr     bool
	}{
		{
			name: "ok",
		},
		{
			name:        "failed: 1",
			argResponse: Response{Failed: 1},
		},
		{
			name:        "failed: 2",
			argResponse: Response{Failed: 2},
		},
		{
			name:        "failed: 3",
			argResponse: Response{Failed: 3},
		},
		{
			name:        "failed: 4",
			argResponse: Response{Failed: 4},
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := lp.checkResponse(tt.argResponse); (err != nil) != tt.wantErr {
				t.Errorf("Longpoll.checkResponse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TODO: write TestLongpoll_Run
// func TestLongpoll_Run(t *testing.T) {
// 	groupToken := os.Getenv("GROUP_TOKEN")
// 	if groupToken == "" {
// 		t.Skip("GROUP_TOKEN empty")
// 	}
// 	vk := api.NewVK(groupToken)
// 	groupID, _ := strconv.Atoi(os.Getenv("GROUP_ID"))
// 	lp, _ := NewLongpoll(&vk, groupID)
// 	lp.Wait = 1

// 	t.Run("Run", func(t *testing.T) {
//
// 	})
// }

func TestLongpoll_RunError(t *testing.T) {
	t.Parallel()

	vk := api.NewVK("")
	lp, _ := Init(vk, 0)
	lp.Wait = 1

	t.Run("Run client error", func(t *testing.T) {
		if err := lp.Run(); err == nil {
			t.Error(err)
		}
	})

	lp.Server = "http://example.com"

	t.Run("Run json error", func(t *testing.T) {
		if err := lp.Run(); err == nil {
			t.Error(err)
		}
	})
}
