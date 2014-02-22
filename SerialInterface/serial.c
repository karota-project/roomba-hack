#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <fcntl.h>
#include <string.h>
#include <termios.h>
#include <time.h>
#include <sys/wait.h>

typedef enum {
    START,
    BAUD,
    CONTROL,
    SAFE,
    FULL,
    POWER,
    SPOT,
    CLEAN,
    MAX,
    DRIVE,
    MOTORS,
    LEDS,
    SONG,
    PLAY,
    SENSORS,
    DOCK  // DOCKへ戻す
} Command;

#define SERIAL_PORT "/dev/ttyUSB0"
#define BAUD_RATE 115200

int CONTINUE = 1;

typedef struct {
    char *name;  // command-name
    int opcode;  // operation-code
    char *data;  // data
} CMD;

extern CMD cmd[] = {
    {"START",128, NULL},  // SCI通信時に必ず行う
    {"BAUD",129, "115200"},  // boud-rate変更 (Data : 1Byte)
    {"CONTROL",130, NULL},  // passive-mode移行
    {"SAFE",131, NULL},  // safa-mode移行
    {"FULL",132, NULL},  // full-mode移行
    {"POWER",133, NULL},  // "POWERボタン"押すのと同様の動作
    {"SPOT",134, NULL},  // "SPOTボタン"(掃除)押すのと同様の動作
    {"CLEAN",135, NULL},  // "CLEANボタン"(掃除)押すのと同様の動作
    {"MAX",136, NULL},  // "MAXボタン"(掃除)押すのと同様の動作
    {"DRIVE",137, ""},  // 4ByteでRoombaを操作する (Velocity : 2byte , Radius : 2byte)
    {"MOTORS",138, ""},  // 1Byte cleaning motors(ブラシ)を操作する
    {"LEDS",139, ""},  // LEDを制御する(Data : 3Byte)
    {"SONG",140, ""},  // 電子音を鳴らす (Data : nByte)
    {"PLAY",141, ""},  // 用意された音楽を鳴らす(Data : 1Byte)
    {"SENSORS",142, ""},  //rommbaにセンサデータを送信? (Data : 1Byte)
    {"DOCK",143, NULL}  // DOCKへ戻す
};

/**
 *  sendCommand
 *
 */
void sendCommand(int fd, int cmdNum) {
    char buf[128];
    memset(&buf, 0, sizeof(buf));

    if (cmd[cmdNum].data != NULL && cmd[cmdNum].data != "") {
        // need space?
        sprintf(buf, "%d%d", cmd[cmdNum].opcode, atoi(cmd[cmdNum].data));

    } else {
        sprintf(buf, "%d", cmd[cmdNum].opcode);
    }

    write(fd, buf, sizeof(buf));

    fprintf(stdout, "Send[%s] Opcode[%d] Data[%s]  BUF[%s]\n",cmd[cmdNum].name, cmd[cmdNum].opcode,cmd[cmdNum].data,buf);
}

/**
 *  writeProcess
 *
 */
void writeProcess(int fd, pid_t result_pid) {
    /* setup */
    // if overwrite rommba_cmd's data
    // cmd[SENSORS].data = "aaaaa";

    sendCommand(fd, START);
    sendCommand(fd, CONTROL);

    /* send */
    char input[256];
    while (CONTINUE != -1) {
        memset(&input, 0, sizeof(input));

        fgets(input, sizeof(input), stdin);

        if (atoi(input) == 99 || atoi(input) == NULL ) {
            CONTINUE = -1;
        } else {
            int num = atoi(input);
            sendCommand(fd,num);
        }
    }
}

/**
 *  readProcess
 *
 */
void readProcess(int fd) {
    char buf[256];
    int ret;

    //STOPになるまで無限ループ
    while (CONTINUE) {
        memset(&buf, 0, sizeof(buf));
        ret = read(fd, &buf, 256);
        if (ret < 0) {
            fprintf(stdout, "Could not read from serial port\n");
            CONTINUE = 0;
        } else {
            fprintf(stdout, "%s\n", buf);
        }
    }

    fprintf(stdout, "readProcess end...\n");
}

/**
 *  initSerial
 *
 */
void initSerial(int fd) {
    struct termios tio;
    memset(&tio,0,sizeof(tio));
    tio.c_cflag = CS8 | CLOCAL | CREAD;
    tio.c_cc[VTIME] = 100;
    // ボーレートの設定
    cfsetispeed(&tio, BAUD_RATE);
    cfsetospeed(&tio, BAUD_RATE);
    // デバイスに設定を行う
    tcsetattr(fd, TCSANOW, &tio);
}

/**
 *  main
 *
 */
void main (int argc, char **argv) {
    printf("option is %s\n", argv[1]);
    printf("option2 is %s\n", argv[2]);

    int fd;

    // シリアルポート準備
    initSerial(fd);

    // forkして受信用と送信用に分ける
    pid_t result_pid = fork();


    if (result_pid == -1) {
        fprintf(stdout, "fork failed.\n");
        return;
    }

    if (result_pid == 0) {
        readProcess(fd); // child
    } else {
        fprintf(stdout, "fork completed.\n");
        writeProcess(fd, result_pid);  // parent
    }

    return;
}

