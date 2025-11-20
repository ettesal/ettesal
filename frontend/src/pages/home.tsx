import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
// import axios from "axios";
import { Greet } from "../../wailsjs/go/main/App";
import { useCallback, useState } from "react";
import { Textarea } from "@/components/ui/textarea";

const Home = () => {

  const [link, setLink] = useState("");
  const [inputData, setInputData] = useState("");



  const clk = useCallback(async () => {
    const t = await Greet(link);
    setInputData(t);
  }, [link]);

  return (
    <div className="container mx-auto my-8 flex gap-3 flex-wrap justify-center">
      <div className="w-full flex gap-3 mx-8">
        <Input
          value={link}
          onChange={(e) => setLink(e.target.value)}
        />

        <Button onClick={clk}>Fetch</Button>

      </div>

      <br />


      <hr />


      <div className="w-10/12 text-wrap">

        <Textarea value={inputData} readOnly/>




        <div className="w-full bg-gray-400">
          <iframe className="h-60 hover:h-screen animate-pulse" src={link} width={'100%'}  ></iframe>
        </div>




      </div>
    </div>
  );
};

export default Home;
