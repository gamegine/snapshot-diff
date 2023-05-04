
# QNAP directory structure

```
/
├──/share
│   ├── <smb share name>: symlink -> <volume x>/<name>
│   └── <volume ex:CE_CACHEDEV1_DATA>
│       └── <share name>
├──/mnt/snapshot/
│   ├── <volume id>
│   |   └── <snapshot id>
│   └── /mnt/snapshot/export/Unified-Snapshot/
│       └── <volume name>
│           └── <snapshot name>: symlink -> /mnt/snapshot/<volume id>/<snapshot id>
```