import { CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { HttpErrorResponse } from '@angular/common/http';
import { RouterTestingModule } from '@angular/router/testing';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MockComponent } from 'ng2-mock-component';
import { StoreModule, Store } from '@ngrx/store';

import { ChefServersListComponent } from './chef-servers-list.component';
import { Server } from 'app/entities/servers/server.model';
import { CreateServerSuccess, CreateServerFailure } from 'app/entities/servers/server.actions';
import { NgrxStateAtom, ngrxReducers, runtimeChecks } from 'app/ngrx.reducers';
import { HttpStatus } from 'app/types/types';
import { FeatureFlagsService } from 'app/services/feature-flags/feature-flags.service';

describe('ChefServersListComponent', () => {
  let component: ChefServersListComponent;
  let fixture: ComponentFixture<ChefServersListComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [
        MockComponent({
        selector: 'app-create-chef-server-modal',
        inputs: ['visible', 'creating', 'conflictErrorEvent', 'createForm'],
        outputs: ['close', 'createClicked']
        }),
        MockComponent({ selector: 'app-delete-object-modal',
        inputs: ['default', 'visible', 'objectNoun', 'objectName'],
        outputs: ['close', 'deleteClicked'] }),
        MockComponent({ selector: 'chef-button',
                inputs: ['disabled', 'routerLink'] }),
        MockComponent({ selector: 'chef-th' }),
        MockComponent({ selector: 'chef-td' }),
        MockComponent({ selector: 'chef-error' }),
        MockComponent({ selector: 'chef-form-field' }),
        MockComponent({ selector: 'chef-heading' }),
        MockComponent({ selector: 'chef-icon' }),
        MockComponent({ selector: 'chef-loading-spinner' }),
        MockComponent({ selector: 'mat-select' }),
        MockComponent({ selector: 'mat-option' }),
        MockComponent({ selector: 'chef-page-header' }),
        MockComponent({ selector: 'chef-subheading' }),
        MockComponent({ selector: 'chef-toolbar' }),
        MockComponent({ selector: 'chef-table-new' }),
        MockComponent({ selector: 'chef-table-header' }),
        MockComponent({ selector: 'chef-table-body' }),
        MockComponent({ selector: 'chef-table-row' }),
        MockComponent({ selector: 'chef-table-header-cell' }),
        MockComponent({ selector: 'chef-table-cell' }),
        MockComponent({ selector: 'a', inputs: ['routerLink'] }),
        ChefServersListComponent
      ],
      providers: [
        FeatureFlagsService
      ],
      imports: [
        FormsModule,
        ReactiveFormsModule,
        RouterTestingModule,
        StoreModule.forRoot(ngrxReducers, { runtimeChecks })
      ],
      schemas: [ CUSTOM_ELEMENTS_SCHEMA ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ChefServersListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  describe('create server', () => {
    let store: Store<NgrxStateAtom>;
    const server = <Server> {
        id: '1',
        name: 'new server',
        description: 'new server description',
        fqdn: 'xyz.com',
        ip_address: '1.1.1.1'
      };

    beforeEach(() => {
      store = TestBed.inject(Store);
    });

    it('openCreateModal opens modal', () => {
      expect(component.createModalVisible).toBe(false);
      component.openCreateModal();
      expect(component.createModalVisible).toBe(true);
    });

    it('opening create modal resets name, description, fqdn and ip_address to empty string', () => {
      component.createChefServerForm.controls['name'].setValue('any');
      component.createChefServerForm.controls['description'].setValue('any');
      component.createChefServerForm.controls['fqdn'].setValue('any');
      component.createChefServerForm.controls['ip_address'].setValue('any');
      component.openCreateModal();
      expect(component.createChefServerForm.controls['name'].value).toBe(null);
      expect(component.createChefServerForm.controls['description'].value).toBe(null);
      expect(component.createChefServerForm.controls['fqdn'].value).toBe(null);
      expect(component.createChefServerForm.controls['ip_address'].value).toBe(null);
    });

    it('on success, closes modal and adds new server', () => {
      component.createChefServerForm.controls['name'].setValue(server.name);
      component.createChefServerForm.controls['description'].setValue(server.description);
      component.createChefServerForm.controls['fqdn'].setValue(server.fqdn);
      component.createChefServerForm.controls['ip_address'].setValue(server.ip_address);
      component.createChefServer();

      store.dispatch(new CreateServerSuccess({'server': server}));

      component.sortedChefServers$.subscribe(servers => {
        expect(servers).toContain(server);
      });
    });

    it('on conflict error, modal is open with conflict error', () => {
      spyOn(component.conflictErrorEvent, 'emit');
      component.openCreateModal();
      component.createChefServerForm.controls['name'].setValue(server.name);
      component.createChefServerForm.controls['description'].setValue(server.description);
      component.createChefServerForm.controls['fqdn'].setValue(server.fqdn);
      component.createChefServerForm.controls['ip_address'].setValue(server.ip_address);
      component.createChefServer();

      const conflict = <HttpErrorResponse>{
        status: HttpStatus.CONFLICT,
        ok: false
      };
      store.dispatch(new CreateServerFailure(conflict));

      expect(component.createModalVisible).toBe(true);
      expect(component.conflictErrorEvent.emit).toHaveBeenCalled();
    });

    it('on create error, modal is closed with failure banner', () => {
      spyOn(component.conflictErrorEvent, 'emit');
      component.openCreateModal();
      component.createChefServerForm.controls['name'].setValue(server.name);
      component.createChefServerForm.controls['description'].setValue(server.description);
      component.createChefServerForm.controls['fqdn'].setValue(server.fqdn);
      component.createChefServerForm.controls['ip_address'].setValue(server.ip_address);
      component.createChefServer();

      const error = <HttpErrorResponse>{
        status: HttpStatus.INTERNAL_SERVER_ERROR,
        ok: false
      };
      store.dispatch(new CreateServerFailure(error));

      expect(component.createModalVisible).toBe(false);
      expect(component.conflictErrorEvent.emit).toHaveBeenCalledWith(false);
    });

  });

  describe('create server form validation', () => {

    it('- fqdn field validity', () => {
      component.openCreateModal();

      let errors = {};
      const fqdn = component.createChefServerForm.controls['fqdn'];
      expect(fqdn.valid).toBeFalsy();

      // FQDN field is required
      errors = fqdn.errors || {};
      expect(errors['required']).toBeTruthy();

      // Set fqdn to invalid inputs
      fqdn.setValue('http://foo.bar-.-.');
      errors = fqdn.errors || {};
      expect(errors['pattern']).toBeTruthy();

      fqdn.setValue('http://foo.bar..com/');
      errors = fqdn.errors || {};
      expect(errors['pattern']).toBeTruthy();

      fqdn.setValue('http://...foo.bar.com/');
      errors = fqdn.errors || {};
      expect(errors['pattern']).toBeTruthy();

      fqdn.setValue('http://...foo.bar com/');
      errors = fqdn.errors || {};
      expect(errors['pattern']).toBeTruthy();

      // Set fqdn to valid inputs
      fqdn.setValue('http://fo_o.bar/');
      errors = fqdn.errors || {};
      expect(errors['pattern']).toBeFalsy();

      fqdn.setValue('https://bit.ly/2OWGwiL');
      errors = fqdn.errors || {};
      expect(errors['pattern']).toBeFalsy();

      fqdn.setValue('http://demo.com/');
      errors = fqdn.errors || {};
      expect(errors['pattern']).toBeFalsy();

      fqdn.setValue('https://example_test.co.in/2OWGwiL');
      errors = fqdn.errors || {};
      expect(errors['pattern']).toBeFalsy();

    });
  });
});

