import potsdb
import datetime
import time

if __name__ == "__main__":

	sample = {
		 u'counter_name': u'floatingip.incoming.packets',
		 u'user_id': u'',
		 u'message_signature': u'86432af2ab889d14061655a534ab478803d7e7708b8a30759c1de461a47bd037',
		 u'timestamp': datetime.datetime(2016,8,26,7,18,29,535244),
		 u'resource_id': u'5ebe4635-9a4b-4870-8eb3-a192ce54b869',
		 u'resource_metadata': u'',
		 u'source': u'openstack',
		 u'counter_unit': u'packets',
		 u'counter_volume': 10,
		 u'project_id': u'13076c4c52c844c6a63dcf3708552c96',
		 u'message_id': u'e803978e-6b54-11e6-b041-40f2e9cefffa',
		 u'counter_type': u'delta'
	}

	print sample['counter_name']
	print sample['timestamp']
	print time.mktime(sample['timestamp'].timetuple())
	print sample['counter_volume']

	metrics = potsdb.Client('127.0.0.1', port=4242, qsize=1, host_tag=True, mps=100, check_host=True)
	try:
		metrics.send(
			sample['counter_name'],
			sample['counter_volume'],
			timestamp=time.mktime(sample['timestamp'].timetuple()),
			resource_id=sample['resource_id'],
			project_id=sample['project_id']
		)
		metrics.send(
			sample['counter_name'],
			sample['counter_volume'],
			timestamp=time.mktime(sample['timestamp'].timetuple()),
			resource_id=sample['resource_id'],
			project_id=sample['project_id']
		)
		metrics.send(
			sample['counter_name'],
			sample['counter_volume'],
			timestamp=time.mktime(sample['timestamp'].timetuple()),
			resource_id=sample['resource_id'],
			project_id=sample['project_id']
		)
	except Exception, e:
		print e
	
	metrics.wait()
