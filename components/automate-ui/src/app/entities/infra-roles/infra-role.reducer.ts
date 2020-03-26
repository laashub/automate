import { EntityState, EntityAdapter, createEntityAdapter } from '@ngrx/entity';
import { set, pipe } from 'lodash/fp';
import { EntityStatus } from 'app/entities/entities';
import { RoleActionTypes, RoleActions } from './infra-role.action';
import { InfraRole } from './infra-role.model';

export interface InfraRoleEntityState extends EntityState<InfraRole> {
  status: EntityStatus;
  getAllStatus: EntityStatus;
}

const GET_ALL_STATUS = 'getAllStatus';

export const infraRoleEntityAdapter: EntityAdapter<InfraRole> = createEntityAdapter<InfraRole>({
  selectId: (infrarole: InfraRole) => infrarole.name
});

export const InfraRoleEntityInitialState: InfraRoleEntityState =
  infraRoleEntityAdapter.getInitialState(<InfraRoleEntityState>{
  getAllStatus: EntityStatus.notLoaded
});

export function infraRoleEntityReducer(
  state: InfraRoleEntityState = InfraRoleEntityInitialState,
  action: RoleActions): InfraRoleEntityState {

  switch (action.type) {
    case RoleActionTypes.GET_ALL:
      return set(GET_ALL_STATUS, EntityStatus.loading, infraRoleEntityAdapter.removeAll(state));

    case RoleActionTypes.GET_ALL_SUCCESS:
      return pipe(
        set(GET_ALL_STATUS, EntityStatus.loadingSuccess))
        (infraRoleEntityAdapter.setAll(action.payload.roles, state)) as InfraRoleEntityState;

    case RoleActionTypes.GET_ALL_FAILURE:
      return set(GET_ALL_STATUS, EntityStatus.loadingFailure, state);

    default:
      return state;
  }
}

export const getEntityById = (name: string) =>
  (state: InfraRoleEntityState) => state.entities[name];
