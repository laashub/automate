BEGIN;

INSERT INTO iam_roles (id, name, type, actions)
  VALUES ('compliance-viewer', 'Compliance Viewer', 'custom', '{compliance:*:get, compliance:*:list}')
  ON CONFLICT (id) DO NOTHING;

INSERT INTO iam_roles (id, name, type, actions)
  VALUES ('compliance-editor', 'Compliance Editor', 'custom', '{compliance:*}')
  ON CONFLICT (id) DO NOTHING;

INSERT INTO iam_policies (id, name, type)
  VALUES ('compliance-access', 'Compliance Viewers', 'custom')
  ON CONFLICT (id) DO NOTHING;

INSERT INTO iam_policies (id, name, type)
  VALUES ('compliance-editor-access', 'Compliance Editors', 'custom')
  ON CONFLICT (id) DO NOTHING;

SELECT insert_iam_statement_into_policy('compliance-access', 'allow', '{}', '{*}', 'compliance-viewer', '{~~ALL-PROJECTS~~}');

SELECT insert_iam_statement_into_policy('compliance-editor-access', 'allow', '{}', '{*}', 'compliance-editor', '{~~ALL-PROJECTS~~}');

END;