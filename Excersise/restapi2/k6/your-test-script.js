import http from 'k6/http';
import { sleep } from 'k6';

export default function () {
  // Send an HTTP GET request to your Go server
  const response = http.post('http://localhost:4002/api/junk1'); // Update the URL

  // Check if the response status code is 200
  if (response.status !== 200) {
    console.error(`Request failed with status code: ${response.status}`);
  }

  // Sleep for 1 second (adjust as needed)
  sleep(1);
}