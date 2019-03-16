from setuptools import setup, find_packages
setup(
	name = "HelloWorld",
	version = "0.0.1",
	packages = find_packages(),
	entry_points = {
		'console_scripts': [
			'lijiao = package1.hello:hello',
		],
		'lijiao.test':[
			'hello = package1.hello:hello',
			'bye   = package1.hello:bye',
		],
	}
)
