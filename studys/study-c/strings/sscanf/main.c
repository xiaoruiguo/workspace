/*
 * main.c
 * Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com>
 *
 * Distributed under terms of the GPL license.
 */

#include <stdio.h>

int main(int argc, char *argv[])
{
	int value;
	int value2;
	int ret;
	ret = sscanf("1000","%d", &value);
	printf("ret is %d, value is %d\n", ret, value);

	ret = sscanf("2000 ","%d", &value);
	printf("ret is %d, value is %d\n", ret, value);

	ret = sscanf(" 3000 ","%d", &value);
	printf("ret is %d, value is %d\n", ret, value);

	ret = sscanf("1000 200 ","%d %d", &value, &value2);
	printf("ret is %d, value is %d\n", ret, value);
	return 0;
}
