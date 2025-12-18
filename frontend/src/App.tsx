import { BrowserRouter, Route, Routes } from "react-router-dom";
import Index from "@/pages/Index";
import NotFound from "@/pages/NotFound";
import { Toaster } from "@/components/ui/toaster";
import { useEffect, useState } from "react";
import { Inicializar } from "../wailsjs/go/main/App";

function App() {
      const [resultText, setResultText] = useState("Please enter your name below 👇");
  
      const updateResultText = (result) => setResultText(result);


      useEffect(() => {
          Inicializar().then((result) => {
              updateResultText(result);
          }).catch((error) => {
              updateResultText(error.toString());
          });
      }, []);
  

  return (
    <BrowserRouter>
      <div className="relative min-h-screen bg-gradient-to-b from-slate-950 via-slate-900 to-slate-950 text-foreground">
        <div className="pointer-events-none absolute inset-0 bg-[radial-gradient(circle_at_20%_20%,rgba(251,191,36,0.08),transparent_35%),radial-gradient(circle_at_80%_0%,rgba(59,130,246,0.06),transparent_30%),radial-gradient(circle_at_50%_80%,rgba(16,185,129,0.06),transparent_30%)]" />
        <main className="relative z-10">
          <Routes>
            <Route path="/" element={<Index />} />
            <Route path="*" element={<NotFound />} />
          </Routes>
        </main>
      </div>
      <Toaster />
    </BrowserRouter>
  );
}

export default App;
