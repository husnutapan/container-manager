import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Observable} from 'rxjs';
import {Namespaces} from '../dto/namespaces';

@Injectable()
export class ContainerService {
  private readonly PATH = `http://localhost:9090`;

  constructor(private http: HttpClient) {
  }

  getContainerNamespaces(): Observable<Namespaces[]> {
    return this.http.get<Namespaces[]>(`${this.PATH}/namespaces`);
  }

  getDatatable(val) {
    return this.http.get<Namespaces[]>(`${this.PATH}/pods/${val}`);

  }

  killPod(selectedNamespaces: string, pod: string) {
    return this.http.get<Namespaces[]>(`${this.PATH}/pods/delete/${selectedNamespaces}/${pod}`);
  }
}
