import {Component, OnInit} from '@angular/core';
import {Namespaces} from '../dto/namespaces';
import {ContainerService} from '../services/container.service';
import {Pod} from '../dto/pod';

@Component({
  templateUrl: './container.component.html',
  selector: 'app-container'
})
export class ContainerComponent implements OnInit {
  namespacesList: Namespaces[];
  selectedNamespaces: Namespaces;
  podList: Pod[] = [];
  datatableRendered = false;

  constructor(private service: ContainerService) {
  }

  ngOnInit(): void {
    this.service.getContainerNamespaces().subscribe((data) => {
      this.namespacesList = (data) as Namespaces[];
    }, undefined);
  }

  getDatatable() {
    this.datatableRendered = false;
    this.service.getDatatable(this.selectedNamespaces.name).subscribe((data) => {
      this.podList = (data) as Pod[];
      this.datatableRendered = true;
    }, undefined);
  }


  kill(val: Pod) {
    this.service.killPod(this.selectedNamespaces.name, val.name).subscribe((data) => {
      this.datatableRendered = false;
    });
  }
}
