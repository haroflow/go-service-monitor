import { useEffect, useState } from 'react'
import './App.css'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faCircle, faWarning } from '@fortawesome/free-solid-svg-icons'
import { ServiceMonitor } from './model/ServiceMonitor';

function App() {
  const [data, setData] = useState<ServiceMonitor | null>(null);
  const [error, setError] = useState("");

  const loadData = () => {
    fetch('/api')
      .then(res => res.json())
      .then(json => {
        json.lastCheckDate = new Date(json.lastCheck);
        setData(json as ServiceMonitor)
        setError("");
      })
      .catch(err => {
        console.error(err);
        setError("Could not load data from server.");
      });
  }

  useEffect(() => {
    loadData();

    const timer = setInterval(() => {
      loadData();
    }, 5000);

    return () => clearInterval(timer);
  }, []);

  return (
    <div className="container max-w-screen-md	mx-auto bg-slate-50
      border border-slate-300 rounded m-3 drop-shadow-lg">
      <Header />
      <div className="p-3">
        {error ? (
          <div className="rounded border bg-red-50 border-red-400 text-red-700 p-3">
            <p>
              <FontAwesomeIcon icon={faWarning} className="mr-2" />
              {error}
            </p>
          </div>
        ) : null}

        {data ? (
          <>
            <p className="mt-2">
              <span className="font-medium">Last check: </span>
              {data.lastCheckDate.toLocaleString() || '-'}
            </p>

            <div className="mt-3 divide-y">
              {data.httpChecks?.map(c =>
                <StatusRow key={`http_${c.displayName}`} displayName={c.displayName} status={c.status} />)}
              {data.tcpChecks?.map(c =>
                <StatusRow key={`tcp_${c.displayName}`} displayName={c.displayName} status={c.status} />)}
              {data.dnsChecks?.map(c =>
                <StatusRow key={`dns_${c.displayName}`} displayName={c.displayName} status={c.status} />)}
            </div>
          </>
        ) : null}
      </div>
    </div>
  )
}

function Header() {
  return (
    <div className="header rounded-t pt-10 p-5 select-none shadow-lg">
      <h1 className="text-white text-4xl font-thin">go-service-monitor</h1>
    </div>
  );
}

function StatusRow(props: {
  displayName: string,
  status: boolean
}) {
  return (
    <div className="p-2 grid grid-cols-2">
      <span>{props.displayName}</span>
      <span><StatusLabel ok={props.status} /></span>
    </div>
  );
}

function StatusLabel(props: { ok: boolean }) {
  return props.ok
    ? <span className="text-green-700">
      <FontAwesomeIcon icon={faCircle} className="drop-shadow" /> OK
    </span>
    : <span className="text-red-600">
      <FontAwesomeIcon icon={faCircle} className="drop-shadow" /> FAILED
    </span>;
}

export default App
