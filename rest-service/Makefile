BASE_URL=localhost:8080
PEOPLE_ROUTE=$(BASE_URL)/people

test:
	# Test initial connection
	nc -z localhost 8080 || (echo "Is the server running on localhost:8080?"; exit 1)
	# Route DNE query -> 404
	curl -sX GET '$(BASE_URL)/person'
	# Incorrect method people query -> 405
	curl -sX POST '$(PEOPLE_ROUTE)'
	# All people query -> 5 results
	curl -sX GET '$(PEOPLE_ROUTE)' | jq
	# All people query with trailing / -> 404 
	curl -sIX GET '$(PEOPLE_ROUTE)/' 
	# People with extra path route DNE -> 404
	curl -sIX GET '$(PEOPLE_ROUTE)/path/doesnt/exist' 
	# Person by Malformed ID query -> 404
	curl -sIX GET '$(PEOPLE_ROUTE)/12341234bd8f9'
	# Person by Non-Existnant ID query -> 404
	curl -sIX GET '$(PEOPLE_ROUTE)/12341234-b659-422b-ab48-a0d0d40bd8f9'
	# Person by ID query -> 1 result
	curl -sX GET '$(PEOPLE_ROUTE)/000ebe58-b659-422b-ab48-a0d0d40bd8f9' | jq
	# Only First Query -> no results
	curl -sX GET '$(PEOPLE_ROUTE)?firstName=John' | jq
	# First and Last dont match Query -> no results
	curl -sX GET '$(PEOPLE_ROUTE)?firstName=John&lastName=Joe' | jq
	# First and Last Query -> 2 results
	curl -sX GET '$(PEOPLE_ROUTE)?firstName=John&lastName=Doe' | jq
	# 2 First and 1 Last Name Query -> 2 results
	curl -sX GET '$(PEOPLE_ROUTE)?firstName=John&lastName=Doe&firstName=Jane' | jq
	# 2 First and 2 Last Names Query -> 3 results
	curl -sX GET '$(PEOPLE_ROUTE)?firstName=John&lastName=Doe&firstName=Jane&lastName=Doe' | jq
	# Phone Number doesnt match Query -> no results
	curl -sX GET '$(PEOPLE_ROUTE)?phoneNumber=+1%20%28800%29%20555-1313' | jq
	# Phone Number Query -> 1 result
	curl -sX GET '$(PEOPLE_ROUTE)?phoneNumber=%2B1%20%28800%29%20555-1313' | jq
	# 2 Phone Numbers Query -> 2 results
	curl -sX GET '$(PEOPLE_ROUTE)?phoneNumber=%2B1%20%28800%29%20555-1313&phoneNumber=%2B1%20%28800%29%20555-1212' | jq
	# Query Param Does Not Exist -> no results
	curl -sX GET '$(PEOPLE_ROUTE)?myRandomQueryParam=doesNothing' | jq

