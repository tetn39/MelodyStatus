import axios from 'axios'
import { useQuery } from "@tanstack/react-query";

interface HelloWorld {
  message: string
}

const fetchHelloWorld = async () => {
  console.info('Fetching...')
  await new Promise((r) => setTimeout(r, 500)) // 疑似的に500msの遅延を挿入
  const response = await axios.get<HelloWorld>('http://localhost:8080/api/v1/hello')
  return response.data
}


export const useHelloWorld = () => {
  return useQuery({
    queryKey: ["HelloWorld"],
    queryFn: fetchHelloWorld,
  });
};
