import http from 'k6/http';

import { sleep } from 'k6';
export const options ={
    vus : 10,
    iterations: 1000000, 
}

export default function () {
    http.post('https://localhost:4000/customer');
     sleep(1);
}