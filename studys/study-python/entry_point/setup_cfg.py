from setuptools import setup, find_packages
setup(
	packages = find_packages(),
	setup_requires=['pbr>=1.8'],
	pbr=True,
)
