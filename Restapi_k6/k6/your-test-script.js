import http from 'k6/http';
import { check } from 'k6';

export default function () {
  const url = 'http://localhost:8080/api'; 
  const requestBody = {
    cusid: 1,
    firstname: "Nithish",
    lastname: "T",
    bankid: "kvb1",
    balance: 55000,
    isactive: true
  };
  const params = {
    headers: {
      'Content-Type': 'application/json',
       Authorization :"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjdXNpZCI6MTIzLCJleHAiOjE2OTQ1MTYxNTR9.uaWbxzUJy3GXeAR4PH0ijRiCHJ6-Nxe-CaPxQ9jj7s0"
    },
  };
  const response = http.post(url, JSON.stringify(requestBody), params);
  check(response, { 'status is 200': (r) => r.status === 200 });
}
