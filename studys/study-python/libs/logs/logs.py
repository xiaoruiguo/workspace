import logging
import logging.handlers

def init_log(logpath,program,level):
    """
    logfile:  logfile will in the absolute directory.
    program:  program name, the logfile will be named program.log
    level:    the threshold for the logger
                  logging.CRITICAL,
                  logging.ERROR,
                  logging.WARNING
                  logging.INFO,
                  logging.DEBUG
                  logging.NOTSET
    """
    logfile = logpath+"/"+program+".log"
    handler= logging.handlers.RotatingFileHandler(logfile,'a',\
            maxBytes=1024*1024*100,backupCount=5)
    fmt = '%(asctime)s %(levelname)s %(filename)s %(lineno)d: %(message)s'
    formatter = logging.Formatter(fmt)
    handler.setFormatter(formatter)

    logger = logging.getLogger(program)
    logger.addHandler(handler)
    logger.setLevel(level)

    return logger

if __name__ == "__main__":
    logger = init_log("/var/log","test",logging.DEBUG)
    logger.info("starting..")

