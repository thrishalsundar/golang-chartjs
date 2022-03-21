import React, { useEffect, useState } from 'react'
import './Linechart.css'
import { Line } from 'react-chartjs-2'
import axios from 'axios'

const Linechart = () => {
  const [items, setitems] = useState([])
  const [isLoading, setisLoading] = useState(true)
  var xlabels=[]
  var ylabels=[]
  const data= {
    labels: ylabels,
    datasets: [{
        label: 'Amount in the account',
        data: xlabels,
        backgroundColor: [
            'rgba(255, 99, 132, 0.2)',
        ],
        borderColor: [
            'rgba(255, 99, 132, 1)',
        ],
        borderWidth: 1
    }]
  }
  
  const options= {
    scales: {
        y: {
            beginAtZero: true
        }
    }
  }
  useEffect(()=>{
    var fetchItems=async()=>{
      const result=await axios(`http://localhost:7999/chartapp/transactions/getall`)
      // console.log(result.data);
      await setitems(result.data)
      await setisLoading(false)
    }

    fetchItems();
  },[])

  items.forEach((item)=>{
    // console.log(item["total"]);
    xlabels.push(item["total"])
    ylabels.push(item["dateOfTrans"])
  })
  

  return (
    <div className="bar_chart">
      <h1>Here we go</h1>
        <div className="chart">
          <Line  options={options} data={data}/>
          </div>
    </div>
  )
}



export default Linechart