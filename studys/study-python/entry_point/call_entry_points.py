import pkg_resources
from pkg_resources import load_entry_point

def run_entry_point():
	group = 'lijiao.test'
	for entrypoint in pkg_resources.iter_entry_points(group=group):
		plugin = entrypoint.load()
		plugin()

if __name__ == "__main__":
	run_entry_point()
	load_entry_point('HelloWorld', 'lijiao.test', 'bye')()
