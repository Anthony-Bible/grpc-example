import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import { createPromiseClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { YourService } from "./gen/connectrpc/";



const transport = createGrpcWebTransport({
    baseUrl: "http://localhost:8080",
});

const client = createPromiseClient(YourService, transport);
function App() {
  return <>Hello World</>
}

export default App
