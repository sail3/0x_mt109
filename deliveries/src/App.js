import React from 'react';
import './style.css';
import routes from './routes.json'

class App extends React.Component {
  state = {
    data: routes.routes
  }

  renderButton(deliverStatus) {
    switch (deliverStatus) {
      case 'ongoing':
        return <button type="button" className="button button-rounded ongoing">{deliverStatus}</button>
        case 'pending': 
        return <button type="button" className="button button-rounded pending">{deliverStatus}</button>
      default:
        break;
    }
  }
  
  renderLineStatus(lineStatus) {
    switch (lineStatus) {
      case 'ongoing':
        return <div className="line-status ongoing">&nbsp;</div>
        case 'pending': 
        return <div className="line-status pending">&nbsp;</div>
      default:
        break;
    }
  }

  render() {
    return (<div>
      <table>
        <tbody>
          {this.state.data.map((route) => (
            <tr>
              <td>{this.renderLineStatus(route.status)}</td>
              <td>
                <label className="table-label">Codigo</label><br/>
                <div className="table-content">{route.id}</div>
              </td>
              <td>
                <label className="table-label">Nombre</label><br/>
                <div className="table-content">{route.driver_name}</div>
              </td>
              <td>
                <label className="table-label">Creada</label><br/>
                <div className="table-content">{route.created_at}</div>
              </td>
              <td>
                <label className="table-label">Entregas</label><br/>
                <div className="table-content">{route.deliveries}</div>
              </td>
              <td>
                {this.renderButton(route.status)}
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>);
  }
}

export default App;