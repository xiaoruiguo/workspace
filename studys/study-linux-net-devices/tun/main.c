#include <stdio.h>
#include <sys/socket.h>
#include <sys/ioctl.h>
#include <getopt.h>
#include <linux/if.h>
#include <stdlib.h>
#include <linux/if_tun.h>
#include <fcntl.h>
#include <string.h>
#include <unistd.h>

int tun_alloc(char *dev)
{
	struct ifreq ifr;
	int fd, err;

	if( (fd = open("/dev/net/tun", O_RDWR)) < 0 ){
		printf("open /dev/net/tun fail\n");
		return -1;
	}

	memset(&ifr, 0, sizeof(ifr));

	/* Flags: IFF_TUN   - TUN device (no Ethernet headers) 
	 *        IFF_TAP   - TAP device  
	 *
	 *        IFF_NO_PI - Do not provide packet information  
	 */ 
	ifr.ifr_flags = IFF_TUN; 
	if( *dev )
		strncpy(ifr.ifr_name, dev, IFNAMSIZ);

	if( (err = ioctl(fd, TUNSETIFF, (void *) &ifr)) < 0 ){
		close(fd);
		return err;
	}
	strcpy(dev, ifr.ifr_name);
	return fd;
}              

int main(int argc, char * const argv[])
{
	int choice;
	char name[20];
	snprintf(name,sizeof(name),"tun-default");
	while (1)
	{
		static struct option long_options[] =
		{
			/* Use flags like so:
			{"verbose",	no_argument,	&verbose_flag, 'V'}*/
			/* Argument styles: no_argument, required_argument, optional_argument */
			{"version", no_argument,	0,	'v'},
			{"help", no_argument,	0,	'h'},
			{"name", required_argument, 0, 'n'},
			{0,0,0,0}
		};

		int option_index = 0;

		/* Argument parameters:
			no_argument: " "
			required_argument: ":"
			optional_argument: "::" */

		choice = getopt_long( argc, argv, "vhn:",
					long_options, &option_index);

		if (choice == -1)
			break;

		switch( choice )
		{
			case 'v':
				printf("version 0.1\n");
				return;

			case 'h':
				printf("%s -n [tun device name]\n", argv[0]);
				return;

			case '?':
				/* getopt_long will have already printed an error */
				break;

			default:
				/* Not sure how to get here... */
				return EXIT_FAILURE;
		}
	}

	int fd = tun_alloc(name);

	if (fd == -1) {
		printf("create error: %d\n", fd);
		return;
	}

	while(1){
		sleep(1000);
	}
}
