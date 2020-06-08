import {BrowserModule} from '@angular/platform-browser';
import {NgModule} from '@angular/core';
import {AppComponent} from './app.component';
import {ContainerComponent} from './home/component/container.component';
import {PanelModule} from 'primeng/panel';
import {ButtonModule, DropdownModule, TableModule} from 'primeng';
import {FormsModule} from '@angular/forms';
import {ContainerService} from './home/services/container.service';
import {HttpClientModule} from '@angular/common/http';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';

@NgModule({
  declarations: [
    AppComponent,
    ContainerComponent
  ],
  imports: [
    BrowserModule,
    PanelModule,
    DropdownModule,
    FormsModule,
    HttpClientModule,
    BrowserAnimationsModule,
    TableModule,
    ButtonModule
  ],
  providers: [ContainerService],
  bootstrap: [AppComponent]
})
export class AppModule {
}
