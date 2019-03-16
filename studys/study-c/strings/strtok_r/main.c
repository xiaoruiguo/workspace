/*
 * main.c
 * Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com>
 *
 * Distributed under terms of the GPL license.
 */

#include <stdio.h>
#include <string.h>

#define str_iterate_parts(__iterator, __splitme, __separators)            \
    for (char *__p = NULL, *__it = strtok_r(__splitme, __separators, &__p); \
            (__iterator = __it);                                            \
            __iterator = __it = strtok_r(NULL, __separators, &__p))

int main(int argc, char *argv[])
{
	char str[100];
	sprintf(str,"a b c d e f");
	printf("before: %s\n", str);
	char *delim = " ";
	char *token = NULL;
	char *saveptr = NULL;

/*
	token = strtok_r(str, delim, &saveptr);
	while (NULL != token) {
		printf("%s\n", token);
		token = strtok_r(NULL, delim, &saveptr);
	}
*/

	int i = 0;
	str_iterate_parts(token, str, delim){
		if (i >= 3 ) break;
		printf("%s\n", token);
		i++;
	}
	printf("after : %s\n", str);

	return 0;
}
